package asynq

import (
	"time"

	"github.com/hibiken/asynq"

	"project-v/internal/pkg/notification"
)

// IProducer ...
type IProducer interface {
	EnqueueAt(
		msg []byte, processAt time.Time,
	) error
}

type IProducerClient interface {
	WithName(topic string) IProducer
}

// AsynqProducer ...
type asynqProducer struct {
	producer  *asynq.Client
	inspector *asynq.Inspector
}

func (p *asynqProducer) Publish(data notification.IDataDelivery) (err error) {
	msg, err := data.Marshal()
	if err != nil {
		return err
	}
	_, err = p.producer.Enqueue(
		asynq.NewTask(data.GetRoutingKey(), msg),
	)
	return err
}

func (p *asynqProducer) PublishWithProcessAt(
	data notification.IDataDelivery, processAt time.Time,
) (err error) {
	msg, err := data.Marshal()
	if err != nil {
		return err
	}
	_, err = p.producer.Enqueue(
		asynq.NewTask(data.GetRoutingKey(), msg),
		asynq.ProcessAt(processAt),
	)
	return err
}

// NewProducer ...
func NewProducer(cfg Config) (
	*asynqProducer, error,
) {
	client := asynq.NewClient(
		asynq.RedisClientOpt{
			Addr: cfg.Redis.Address, Password: cfg.Redis.Password,
			DB: cfg.Redis.Database,
		},
	)
	defer client.Close()

	inspector := asynq.NewInspector(
		asynq.RedisClientOpt{
			Addr: cfg.Redis.Address, Password: cfg.Redis.Password,
			DB: cfg.Redis.Database,
		},
	)
	defer inspector.Close()

	return &asynqProducer{
		client,
		inspector,
	}, nil
}

// WithName ...
func (p *asynqProducer) WithName(name string) IProducer {
	return &namedProducer{p, name}
}

type namedProducer struct {
	*asynqProducer

	name string
}

func (p namedProducer) EnqueueAt(
	msg []byte, processAt time.Time,
) error {
	_, err := p.producer.Enqueue(
		asynq.NewTask(p.name, msg), asynq.ProcessAt(processAt),
	)
	return err
}

func (p namedProducer) DeleteTask(id string) error {
	return p.inspector.DeleteTask(p.name, id)
}
