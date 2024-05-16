package rabbitmq

import (
	"errors"
)

var (
	RabbitProducerTimeout   = 10000
	ErrSendToClosedProducer = errors.New("send to closed producer...exiting")
	_defaultMaxThread       = 1
)

const (
	_defaultQueueSize           = 1000
	_defaultNetworkTimeoutInSec = 25
	_defaultRetries             = 10
)

//// RmqConfig ...
//type RmqConfig struct {
//	Name       string                 `json:"name" mapstructure:"name"`
//	Type       string                 `json:"type,omitempty" mapstructure:"type,omitempty"`
//	Exclusive  bool                   `json:"exclusive,omitempty" mapstructure:"exclusive,omitempty"`
//	Durable    bool                   `json:"durable,omitempty" mapstructure:"durable,omitempty"`
//	AutoDelete bool                   `json:"auto_delete,omitempty" mapstructure:"auto_delete,omitempty"`
//	Internal   bool                   `json:"internal,omitempty" mapstructure:"internal,omitempty"`
//	NoWait     bool                   `json:"no_wait,omitempty" mapstructure:"no_wait,omitempty"`
//	Others     map[string]interface{} `json:"others,omitempty" mapstructure:"others,omitempty"`
//	RoutingKey string                 `json:"routing_key" mapstructure:"routing_key"`
//}

// QueueConfig - config queue for consumer
type QueueConfig struct {
	Name       string                 `json:"name" mapstructure:"name"` //require
	Durable    bool                   `json:"durable" mapstructure:"durable"`
	AutoDelete bool                   `json:"auto_delete" mapstructure:"auto_delete"`
	Exclusive  bool                   `json:"exclusive" mapstructure:"exclusive"`
	NoWait     bool                   `json:"no_wait" mapstructure:"no_wait"`
	Args       map[string]interface{} `json:"args" mapstructure:"args"`
}

// QueueBindingConfig - config binding for queue consumer
type QueueBindingConfig struct {
	Key      string                 `json:"key" mapstructure:"key"`           //require
	Exchange string                 `json:"exchange" mapstructure:"exchange"` //require
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

// RabbitMqConfiguration ...
type RabbitMqConfiguration struct {
	// amqp://user:pass@host:port/vhost?heartbeat=10&connection_timeout=10000&channel_max=100
	URI                string             `json:"uri" mapstructure:"uri"`
	QueueBindingConfig QueueBindingConfig `json:"queue_binding_config" mapstructure:"queue_binding_config"`
	QueueConfig        QueueConfig        `json:"queue_config" mapstructure:"queue_config"`
	ConsumeConfig      ConsumeConfig      `json:"consume_config" mapstructure:"consume_config"`
	Retries            int                `json:"retries" mapstructure:"retries"`                         // the number of retries connect to rabbitmq
	InternalQueueSize  int                `json:"internal_queue_size" mapstructure:"internal_queue_size"` // default = 100
	MaxThread          int                `json:"max_thread" mapstructure:"max_thread"`
}
