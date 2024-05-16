package asynq

import (
	"context"
	"log"
	"time"

	"github.com/hibiken/asynq"
)

// IConsumer ...
type IConsumer interface {
	Handle(
		name string, handler EventHandler,
	)
	Start()
}

// EventHandler ...
type EventHandler func(context.Context, *asynq.Task) error

type consumer struct {
	mux *asynq.ServeMux
	cfg Config
}

func NewAsynqConsumer(cfg Config) (*consumer, error) {
	mux := asynq.NewServeMux()
	mux.Use(loggingMiddleware)
	return &consumer{
		mux: mux,
		cfg: cfg,
	}, nil
}

// Handle ...
func (c *consumer) Handle(
	name string, handler EventHandler,
) {
	c.mux.HandleFunc(name, handler)
}

func loggingMiddleware(h asynq.Handler) asynq.Handler {
	return asynq.HandlerFunc(
		func(ctx context.Context, t *asynq.Task) error {
			start := time.Now()
			log.Printf("Start processing %q", t.Type())
			err := h.ProcessTask(ctx, t)
			if err != nil {
				return err
			}
			log.Printf(
				"Finished processing %q: Elapsed Time = %v", t.Type(),
				time.Since(start),
			)
			return nil
		},
	)
}

// Start ...
func (c *consumer) Start() {
	srv := asynq.NewServer(
		asynq.RedisClientOpt{
			Addr: c.cfg.Redis.Address, Password: c.cfg.Redis.Password,
			DB: c.cfg.Redis.Database,
		},
		asynq.Config{
			Concurrency: c.cfg.Concurrency,
		},
	)
	if err := srv.Run(c.mux); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}
