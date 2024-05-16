package event

import (
	"fmt"
	"reflect"

	"project-v/internal/pkg/notification"
)

// eventDispatcherImpl implement IEventDispatcher
type eventDispatcherImpl struct {
	rabbitmqClient notification.IProducer

	// eventBusClient
	// eventBusRedisClient notification.IProducer
	asynqClient notification.IProducer

	mapFactory map[string]notification.IEventFactory
}

// NewDispatcher ...
func NewDispatcher() *eventDispatcherImpl {
	return &eventDispatcherImpl{
		mapFactory: make(map[string]notification.IEventFactory),
	}
}

func (o *eventDispatcherImpl) WithRabbitmqClient(rabbitmqClient notification.IProducer) {
	o.rabbitmqClient = rabbitmqClient
}
func (o *eventDispatcherImpl) WithAsynqClient(asynqClient notification.IProducer) {
	o.asynqClient = asynqClient
}

func (o eventDispatcherImpl) getEventFac(ev notification.INotification) notification.IEventFactory {
	var obj notification.IEventFactory
	var ok bool

	obj, ok = o.mapFactory[reflect.TypeOf(ev).String()]
	if ok {
		return obj
	}

	switch ev.(type) {
	default:
		fmt.Printf(
			"[eventDispatcherImpl] not supported: %v \n",
			reflect.TypeOf(ev).String(),
		)
		return nil
	}
}

func (o eventDispatcherImpl) AddEvent(ev notification.INotification) {
	obj := o.getEventFac(ev)
	if obj == nil {
		return
	}
	obj.Publish(ev)
}
