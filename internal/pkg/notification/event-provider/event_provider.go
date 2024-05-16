package eventprovider

const (
	RabbitMQ      = "rabbitmq"
	EventBus      = "event_bus"
	EventBusRedis = "event_bus_redis"
	Asynq         = "asynq"
)

// MapProviders define the provider available
// if not config here then all event will not able use that provider
var MapProviders = map[string]struct{}{
	RabbitMQ:      {},
	EventBus:      {},
	EventBusRedis: {},
	Asynq:         {},
}
