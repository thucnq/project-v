package appworker

import (
	"context"
	"project-v/config"
	"project-v/pkg/l"
	"project-v/pkg/worker"
)

type Worker struct {
	ctx context.Context
	cfg *config.Config
	w   worker.IWorker

	Log l.Logger `di:"inject"`
}

func New(ctx context.Context, cfg *config.Config) *Worker {
	return &Worker{
		ctx: ctx,
		cfg: cfg,
	}
}

func (w *Worker) Register() {
	builder := worker.New(w.ctx, w.Log)

	builder.WithRabbitmqConsumerConfig(
		&worker.ConsumerRabbitmqGroupConfig{
			URI:               w.cfg.RabbitmqConfig.URI,
			Retries:           w.cfg.RabbitmqConfig.Retries,
			InternalQueueSize: w.cfg.RabbitmqConfig.InternalQueueSize,
			MaxThread:         w.cfg.RabbitmqConfig.MaxThread,
		}, w.cfg.RabbitmqConfig.Events,
	)

	builder.InitSchedule(w.cfg.ScheduleConfig)
	builder.BuildConsumerConfig()
	w.w = builder.Build()
}

// Start ...
func (w Worker) Start() {
	w.w.Start()
}
