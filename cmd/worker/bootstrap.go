package main

import (
	"context"
	appworker "project-v/internal/app-worker"
	handleossignal "project-v/pkg/handle-os-signal"
	"time"

	"project-v/config"
	"project-v/internal/app-api/handler/health"
	"project-v/internal/model/event"
	"project-v/internal/pkg/notification"
	"project-v/pkg/broker/asynq"
	"project-v/pkg/broker/rabbitmq"
	"project-v/pkg/container"
	"project-v/pkg/l"
)

func bootstrap(cfg *config.Config) {
	var log = container.ResolverMust[l.Logger]()

	var shutdown = container.ResolverMust[handleossignal.IShutdownHandler]()

	loadI18n()

	container.Register(
		func() health.Controller {
			return health.New()
		},
	)

	ctx, cancel := context.WithCancel(context.Background())
	shutdown.HandleDefer(cancel)

	// region register Repository - redis
	// endregion

	// region register Repository - db
	// - init db connection
	//_ = mysqlconnect.NewConnectMysql(&cfg.MySql)
	// endregion

	// region register Repository - agency
	// endregion

	// region register Repository - grpc
	// endregion

	// region even bus
	if err := loadEvents(cfg); err != nil {
		log.Error("failed to load events", l.Error(err))
		time.Sleep(5 * time.Second)
		log.Fatal("force shutdown!")
		return
	}
	// endregion

	// region register factory
	// endregion

	// region register publisher
	rabbitmqProducer := rabbitmq.New(
		&rabbitmq.RabbitMqConfiguration{
			URI:       cfg.RabbitmqConfig.URI,
			Retries:   cfg.RabbitmqConfig.Retries,
			MaxThread: cfg.RabbitmqConfig.MaxThread,
		}, log,
	)
	rabbitmqProducer.Start()
	asynqProducer, _ := asynq.NewProducer(
		asynq.Config{
			Redis: asynq.Redis{
				Address:  cfg.Redis.Address,
				Database: cfg.Redis.Database,
				Password: cfg.Redis.Password,
			},
			Concurrency: cfg.ScheduleConfig.Asynq.Concurrency,
		},
	)

	dispatcherObj := event.NewDispatcher()
	dispatcherObj.WithRabbitmqClient(rabbitmqProducer)
	dispatcherObj.WithAsynqClient(asynqProducer)
	container.Register[notification.IEventDispatcher](dispatcherObj)

	// region register service to container
	// endregion

	workerInstance := appworker.New(ctx, cfg)
	container.Fill(workerInstance)

	workerInstance.Register()
	workerInstance.Start()
}
