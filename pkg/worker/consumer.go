package worker

import (
	"context"
	"encoding/json"
	"project-v/pkg/worker/process-status-code"
	"runtime"
	"sync"
	"time"

	"project-v/pkg/broker/rabbitmq"
	"project-v/pkg/l"

	"github.com/streadway/amqp"
)

var stackSize int = 4 << 10 // 4 KB

// NewConsumer ...
func NewConsumer(
	ctx context.Context, ll l.Logger, consumerGroup []*ConsumerGroup,
) *Consumer {
	return &Consumer{
		ctx:         ctx,
		lock:        sync.Mutex{},
		ll:          ll,
		ConsumerJob: consumerGroup,
	}
}

// Close ...
func (s *Consumer) Close() {
	s.ll.Info("Consumer closed")
}

// Start ...
func (s *Consumer) Start() {
	if len(s.ConsumerJob) == 0 {
		return
	}
	s.ll.Info("Start Consumer")

	wait := &sync.WaitGroup{}
	for _, item := range s.ConsumerJob {
		for _, consumerJobCfg := range item.handlerConfigs {
			// todo: check mode and start consumer
			switch item.mode {
			case modeRabbitmq:
				groupConfig := item.cfg.(*ConsumerRabbitmqGroupConfig)
				cfg := &rabbitmq.RabbitMqConfiguration{
					URI: groupConfig.URI,
					QueueBindingConfig: rabbitmq.QueueBindingConfig{
						Key:      consumerJobCfg.QueueBindingConfig.Key,
						Exchange: consumerJobCfg.QueueBindingConfig.Exchange,
						Type:     consumerJobCfg.QueueBindingConfig.Type,
						Name:     consumerJobCfg.QueueBindingConfig.Name,
						NoWait:   consumerJobCfg.QueueBindingConfig.NoWait,
						Args:     consumerJobCfg.QueueBindingConfig.Args,
					},
					QueueConfig: rabbitmq.QueueConfig{
						Name:       consumerJobCfg.QueueConfig.Name,
						Durable:    consumerJobCfg.QueueConfig.Durable,
						AutoDelete: consumerJobCfg.QueueConfig.AutoDelete,
						Exclusive:  consumerJobCfg.QueueConfig.Exclusive,
						NoWait:     consumerJobCfg.QueueConfig.NoWait,
						Args:       consumerJobCfg.QueueConfig.Args,
					},
					ConsumeConfig: rabbitmq.ConsumeConfig{
						Consumer:  consumerJobCfg.ConsumeConfig.Consumer,
						AutoAck:   consumerJobCfg.ConsumeConfig.AutoAck,
						Exclusive: consumerJobCfg.ConsumeConfig.Exclusive,
						NoLocal:   consumerJobCfg.ConsumeConfig.NoLocal,
						NoWait:    consumerJobCfg.ConsumeConfig.NoWait,
						Args:      consumerJobCfg.ConsumeConfig.Args,
					},
					Retries:           groupConfig.Retries,
					InternalQueueSize: groupConfig.InternalQueueSize,
					MaxThread:         groupConfig.MaxThread,
				}
				c := rabbitmq.NewConsumer(s.ctx, cfg, s.ll)
				var loop = 0
				var err error
				for {
					loop++
					err = c.Connect()
					if err != nil {
						if loop > cfg.Retries {
							s.ll.Fatal(
								"failed to connect to rabbitmq", l.Error(err),
							)
							break
						}
						time.Sleep(30 * time.Second)
						continue
					}
					break
				}

				d, err := c.SetupQueue()
				if err != nil {
					s.ll.Fatal("failed to setup queue", l.Error(err))
					break
				}

				wait.Add(1)
				go s.startConsumerRabbitmq(c, d, cfg, consumerJobCfg, wait)
				break
			default:
				s.ll.Warn("unknown mode", l.Any("mode", item.mode))
			}
		}
	}
	wait.Wait()
	s.ll.Info("Exist consumers")
}

func (s *Consumer) startConsumerRabbitmq(
	c *rabbitmq.Consumer, d <-chan amqp.Delivery,
	cfg *rabbitmq.RabbitMqConfiguration, consumerJobCfg *ConsumerHandlerConfig,
	wait *sync.WaitGroup,
) {
	defer wait.Done()
	c.Handle(
		d, func(deliveries <-chan amqp.Delivery) {
			for {
				select {
				case delivery := <-deliveries:
					func(delivery amqp.Delivery) {
						defer func() {
							if r := recover(); r != nil {
								stack := make([]byte, stackSize)
								length := runtime.Stack(stack, true)
								s.ll.Error(
									"have a panic when process message",
									l.String("err", string(stack[:length])),
									l.String(
										"queue",
										consumerJobCfg.QueueConfig.Name,
									),
									l.String(
										"exchange",
										consumerJobCfg.QueueBindingConfig.Exchange,
									),
									l.String("data", string(delivery.Body)),
								)
							}
						}()
						if len(delivery.Body) > 0 {
							event := &Event{}
							err := json.Unmarshal(delivery.Body, event)
							if err != nil {
								s.ll.Error(
									"failed to parse msg", l.Error(err),
									l.String(
										"Received message",
										string(delivery.Body),
									),
								)
								return
							}
							if event.ExchangeName != cfg.QueueBindingConfig.Exchange || event.RoutingKey != cfg.QueueBindingConfig.Key {
								return
							}
							resp := consumerJobCfg.Handler.Handle(delivery.Body)
							switch resp.Code {
							case processstatuscode.Success:
								// atomic.AddInt64(&success, 1)
								if !cfg.ConsumeConfig.AutoAck {
									_ = delivery.Ack(false)
								}
							case processstatuscode.Retry:
								_ = delivery.Nack(false, true)
								break
							case processstatuscode.Drop:
								// atomic.AddInt64(&drop, 1)
								if !cfg.ConsumeConfig.AutoAck {
									_ = delivery.Ack(false)
								}
								// todo: with retry status. push back to queue
							default:
								// atomic.AddInt64(&drop, 1)
								if !cfg.ConsumeConfig.AutoAck {
									_ = delivery.Ack(false)
								}
							}
						}
					}(delivery)

				case <-s.ctx.Done():
					c.Close()
					s.ll.S.Infof(
						"Exiting ... %v.%v", cfg.QueueBindingConfig.Exchange,
						cfg.QueueConfig.Name,
					)
					return
				default:
					time.Sleep(100 * time.Millisecond)
				}
			}
		}, cfg.MaxThread,
	)
}
