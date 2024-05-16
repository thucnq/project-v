package rabbitmq

import (
	"context"
	"errors"
	"fmt"
	"net"
	"strings"
	"time"

	"project-v/pkg/l"

	"github.com/streadway/amqp"
)

// Consumer ...
type Consumer struct {
	ctx context.Context
	ll  l.Logger

	conn    *amqp.Connection
	channel *amqp.Channel
	done    chan error
	cfg     *RabbitMqConfiguration
	retries int // the number of retries to connect to rabbitmq
}

// NewConsumer ...
func NewConsumer(ctx context.Context, cfg *RabbitMqConfiguration, ll l.Logger) *Consumer {
	if cfg == nil {
		return nil
	}
	if cfg.InternalQueueSize == 0 {
		cfg.InternalQueueSize = _defaultQueueSize
	}

	if cfg.MaxThread == 0 {
		cfg.MaxThread = _defaultMaxThread
	}
	return &Consumer{
		ctx:     ctx,
		ll:      ll,
		done:    make(chan error),
		cfg:     cfg,
		retries: _defaultRetries,
	}
}

func (c *Consumer) Connect() error {
	var err error
	// for {
	c.ll.S.Infof("Connecting to rabbitmq on %s", c.cfg.URI)

	c.conn, err = amqp.DialConfig(
		c.cfg.URI, amqp.Config{
			Dial: func(network, addr string) (net.Conn, error) {
				return net.DialTimeout(network, addr, _defaultNetworkTimeoutInSec*time.Second)
			},
		},
	)
	if err != nil {
		return fmt.Errorf("Dial: %s", err)
	}
	go func() {
		// Waits here for the channel to be closed
		c.ll.S.Debugf("Closing: %s", <-c.conn.NotifyClose(make(chan *amqp.Error)))
		// Let Handle know it's not time to reconnect
		c.done <- errors.New("Channel Closed")
	}()

	c.ll.Info("got Connection, getting Channel")
	c.channel, err = c.conn.Channel()
	if err != nil {
		return fmt.Errorf("Channel: %s", err)
	}

	c.ll.S.Infof("got Channel, declaring Exchange (%q)", c.cfg.QueueBindingConfig.Exchange)
	if err = c.channel.ExchangeDeclare(
		c.cfg.QueueBindingConfig.Exchange,
		c.cfg.QueueBindingConfig.Type,
		c.cfg.QueueConfig.Durable,
		c.cfg.QueueConfig.AutoDelete,
		false,
		c.cfg.QueueBindingConfig.NoWait, // noWait
		c.cfg.QueueBindingConfig.Args,   // arguments
	); err != nil {
		return fmt.Errorf("Exchange Declare: %s", err)
	}

	return nil
	// time.Sleep(delay * time.Second)
	// }
}

func (c *Consumer) SetupQueue() (<-chan amqp.Delivery, error) {
	queue, err := c.channel.QueueDeclare(
		c.cfg.QueueConfig.Name,
		c.cfg.QueueConfig.Durable,
		c.cfg.QueueConfig.AutoDelete,
		c.cfg.QueueConfig.Exclusive,
		c.cfg.QueueConfig.NoWait,
		c.cfg.QueueConfig.Args,
	)
	if err != nil {
		return nil, fmt.Errorf("Queue Declare: %s", err)
	}

	c.ll.S.Infof(
		"declared Queue (%q %d messages, %d consumers), binding to Exchange (key %q)",
		queue.Name, queue.Messages, queue.Consumers, c.cfg.QueueBindingConfig.Key,
	)
	if err = c.channel.QueueBind(
		c.cfg.QueueConfig.Name,
		c.cfg.QueueBindingConfig.Key,
		c.cfg.QueueBindingConfig.Exchange,
		c.cfg.QueueBindingConfig.NoWait,
		c.cfg.QueueBindingConfig.Args,
	); err != nil {
		return nil, fmt.Errorf("Queue Bind: %s", err)
	}

	deliveries, err := c.channel.Consume(
		c.cfg.QueueConfig.Name,
		c.cfg.ConsumeConfig.Consumer,
		c.cfg.ConsumeConfig.AutoAck,
		c.cfg.ConsumeConfig.Exclusive,
		c.cfg.ConsumeConfig.NoLocal,
		c.cfg.ConsumeConfig.NoWait,
		c.cfg.ConsumeConfig.Args,
	)

	if err != nil {
		return nil, fmt.Errorf("Queue Consume: %s", err)
	}

	return deliveries, nil
}

func (c *Consumer) Handle(d <-chan amqp.Delivery, fn func(<-chan amqp.Delivery), threads int) {
	var err error
	c.ll.Debug("Consumer Handle START")

	for {
		for i := 0; i < threads; i++ {
			go fn(d)
		}

		// Go into reconnect loop when
		// c.done is passed non nil values
		if e := <-c.done; e != nil {
			if strings.Contains(e.Error(), "Channel Closed") { // retry
				d, err = c.reconnect()
				retries := 0
				for err != nil {

					// Very likely chance of failing
					// should not cause worker to terminate
					retries++
					if retries > c.retries {
						c.ll.Fatal("Cannot reconnect to rabbitmq")
					}
					d, err = c.reconnect()
				}
			} else { // stop
				return
			}
		}
	}
}

func (c *Consumer) reconnect() (<-chan amqp.Delivery, error) {
	c.ll.Info("Consumer - reconnect")
	time.Sleep(30 * time.Second)

	if err := c.Connect(); err != nil {
		return nil, err
	}

	deliveries, err := c.SetupQueue()
	if err != nil {
		return deliveries, errors.New("couldn't connect")
	}

	return deliveries, nil
}

// Close ...
func (c *Consumer) Close() {
	c.done <- errors.New("Stop Consumer")
	_ = c.channel.Close()
	_ = c.conn.Close()
}
