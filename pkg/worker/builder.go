package worker

import (
	"context"
	"fmt"

	aq "github.com/hibiken/asynq"

	"project-v/pkg/l"
)

const (
	modeRabbitmq = "rabbitmq"
)

// Builder ...
type Builder interface {
	Build() IWorker

	WithScheduleHandler(cfg ScheduleHandlerConfig) Builder
	WithScheduleHandle(
		name string, handler func(ctx context.Context, t *aq.Task) error,
	) Builder
	InitSchedule(cfg ScheduleConfig) Builder

	WithRabbitmqConsumerConfig(
		cfg *ConsumerRabbitmqGroupConfig, chConfigs []*ConsumerHandlerConfig,
	) Builder
	WithRabbitmqConsumerHandler(key string, h IConsumerTask) Builder

	BuildConsumerConfig() Builder
}

type IConsumerGroupConfig interface{}

// ConsumerAsynqGroupConfig ...IConsumerGroupConfig
type ConsumerAsynqGroupConfig struct {
	RedisAddr   string `json:"redis_addr" mapstructure:"redis_addr"`
	Concurrency int    `json:"concurrency" mapstructure:"concurrency"`
}

// ConsumerRabbitmqGroupConfig ...IConsumerGroupConfig
type ConsumerRabbitmqGroupConfig struct {
	URI               string `json:"uri" mapstructure:"uri"`
	Retries           int    `json:"retries" mapstructure:"retries"`                         // the number of retries connect to rabbitmq
	InternalQueueSize int    `json:"internal_queue_size" mapstructure:"internal_queue_size"` // default = 100
	MaxThread         int    `json:"max_thread" mapstructure:"max_thread"`
}

// QueueConfig - config queue for consumer
type QueueConfig struct {
	Name       string                 `json:"name" mapstructure:"name"` // require
	Durable    bool                   `json:"durable" mapstructure:"durable"`
	AutoDelete bool                   `json:"auto_delete" mapstructure:"auto_delete"`
	Exclusive  bool                   `json:"exclusive" mapstructure:"exclusive"`
	NoWait     bool                   `json:"no_wait" mapstructure:"no_wait"`
	Args       map[string]interface{} `json:"args" mapstructure:"args"`
}

// QueueBindingConfig - config binding for queue consumer
type QueueBindingConfig struct {
	Key      string                 `json:"key" mapstructure:"key"`           // require
	Exchange string                 `json:"exchange" mapstructure:"exchange"` // require
	Name     string                 `json:"name" mapstructure:"name"`
	NoWait   bool                   `json:"no_wait" mapstructure:"no_wait"`
	Args     map[string]interface{} `json:"args" mapstructure:"args"`
	Type     string                 `json:"type" mapstructure:"type"` // direct|fanout|topic|x-custom
}

// ConsumeConfig - config consume for consumer
type ConsumeConfig struct {
	Consumer  string                 `json:"consumer" mapstructure:"consumer"`
	AutoAck   bool                   `json:"auto_ack" mapstructure:"auto_ack"`
	Exclusive bool                   `json:"exclusive" mapstructure:"exclusive"`
	NoLocal   bool                   `json:"no_local" mapstructure:"no_local"`
	NoWait    bool                   `json:"no_wait" mapstructure:"no_wait"`
	Args      map[string]interface{} `json:"args" mapstructure:"args"`
}

// ConsumerHandlerConfig ...
type ConsumerHandlerConfig struct {
	// config for rabbitmq
	QueueBindingConfig QueueBindingConfig `json:"queue_binding_config" mapstructure:"queue_binding_config"`
	QueueConfig        QueueConfig        `json:"queue_config" mapstructure:"queue_config"`
	ConsumeConfig      ConsumeConfig      `json:"consume_config" mapstructure:"consume_config"`

	GroupID string `json:"group_id" mapstructure:"group_id"`
	Topic   string `json:"topic" mapstructure:"topic"`

	// config general
	Handler IConsumerTask `json:"handler" mapstructure:"handler"`

	mode string // pass from ConsumerGroup
}

func (o ConsumerHandlerConfig) GetKey() string {
	// todo: gen key for rabbitmq
	switch o.mode {
	case modeRabbitmq:
		return fmt.Sprintf(
			"%v.%v", o.QueueBindingConfig.Exchange, o.QueueConfig.Name,
		)
	default:
		panic(fmt.Errorf("unknown mode: %v", o.mode))
	}
}

// ConsumerGroup ...
type ConsumerGroup struct {
	ll             l.Logger
	cfg            IConsumerGroupConfig
	mode           string // rabbitmq
	handlerConfigs map[string]*ConsumerHandlerConfig
}

type builder struct {
	ll        l.Logger
	ctx       context.Context
	scheduler *Scheduler
	// store prepare config to init schedule
	mapScheduleHandler map[string]*ScheduleJobOption

	mapConsumerHandler map[string]map[string]IConsumerTask

	// store prepare config to init consumer
	consumerGroup []*ConsumerGroup

	consumer *Consumer
}

// ScheduleHandlerConfig ...
type ScheduleHandlerConfig struct {
	Handler ISchedulerTask `json:"handler" mapstructure:"handler"`
	Retries int            `json:"retries" mapstructure:"retries"`
	Spec    string         `json:"spec" mapstructure:"spec"`
	Type    string         `json:"type" mapstructure:"type"` // can be: once ,	endless
}

// New ...
func New(ctx context.Context, ll l.Logger) Builder {
	return &builder{
		ctx:                ctx,
		ll:                 ll,
		mapScheduleHandler: make(map[string]*ScheduleJobOption),
		mapConsumerHandler: make(map[string]map[string]IConsumerTask),
	}
}

func (b *builder) Build() IWorker {
	w := &worker{
		scheduler: b.scheduler,
		consumer:  b.consumer,
	}
	return w
}

// WithScheduleHandler ...
func (b *builder) WithScheduleHandler(cfg ScheduleHandlerConfig) Builder {
	if cfg.Handler == nil {
		b.ll.Fatal("missing handler")
	}
	b.mapScheduleHandler[cfg.Handler.GetName()] = &ScheduleJobOption{
		spec:         cfg.Spec,
		pl:           cfg.Handler,
		retries:      cfg.Retries,
		scheduleType: cfg.Type,
		name:         cfg.Handler.GetName(),
	}
	return b
}

// WithScheduleHandle ...
func (b *builder) WithScheduleHandle(
	name string, handler func(ctx context.Context, t *aq.Task) error,
) Builder {
	b.scheduler.WithHandler(name, handler)
	return b
}

// InitSchedule ...
func (b *builder) InitSchedule(cfg ScheduleConfig) Builder {
	var sj []*ScheduleJobOption
	for _, item := range b.mapScheduleHandler {
		sj = append(sj, item)
	}
	// if len(sj) == 0 {
	// 	return b
	// }
	b.scheduler = NewScheduler(b.ctx, sj, b.ll, cfg)
	return b
}

func (b *builder) WithRabbitmqConsumerConfig(
	cfg *ConsumerRabbitmqGroupConfig, chConfigs []*ConsumerHandlerConfig,
) Builder {
	g := &ConsumerGroup{
		cfg:            cfg,
		ll:             b.ll,
		mode:           modeRabbitmq,
		handlerConfigs: make(map[string]*ConsumerHandlerConfig),
	}
	for _, val := range chConfigs {
		val.mode = modeRabbitmq
		g.handlerConfigs[val.GetKey()] = val
	}
	// init consumerGroup with group and topic but not have handler
	b.consumerGroup = append(b.consumerGroup, g)
	return b
}

// WithRabbitmqConsumerHandler ...attach handler by key
func (b *builder) WithRabbitmqConsumerHandler(
	key string, h IConsumerTask,
) Builder {
	if b.mapConsumerHandler[modeRabbitmq] == nil {
		b.mapConsumerHandler[modeRabbitmq] = make(map[string]IConsumerTask)
	}
	b.mapConsumerHandler[modeRabbitmq][key] = h
	return b
}

// BuildConsumerConfig ... map handler to config
func (b *builder) BuildConsumerConfig() Builder {

	if len(b.consumerGroup) == 0 {
		return b
	}

	if b.mapConsumerHandler != nil {
		for _, c := range b.consumerGroup { // list config group
			for key, handlerCfg := range c.handlerConfigs { // list config handle but not have handler yet. need to mapping here
				handler, found := b.mapConsumerHandler[c.mode][key]
				if !found {
					b.ll.Fatal(
						"not found handler for queue",
						l.String("key", key),
					)
				}
				handlerCfg.Handler = handler
			}
		}
	}
	b.consumer = NewConsumer(b.ctx, b.ll, b.consumerGroup)

	return b
}
