package notification

import (
	"time"
)

//go:generate mockgen -source=index.go -destination ./index_mock.go -package=notification
type INotification interface {
	GetData() []byte // convert data to byte array
	GetKey() string
}
type IDataDelivery interface {
	GetExchangeName() string
	GetRoutingKey() string
	Marshal() ([]byte, error)
}
type INotificationVisitor interface {
	Send(INotification)
	SendAt(INotification, time.Time)
}

// IProducer ...
type IProducer interface {
	Publish(data IDataDelivery) (err error)
	PublishWithProcessAt(data IDataDelivery, processAt time.Time) (err error)
}

// IEventFactory ...
type IEventFactory interface {
	Publish(noti INotification)
	WithRabbitmq(rabbitmqClient IProducer)
	WithEventBusRedis(redisClient IProducer)
	WithAsynq(asynqClient IProducer)
}

type IEventDispatcher interface {
	AddEvent(INotification)
}
