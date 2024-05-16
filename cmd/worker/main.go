package main

import (
	"project-v/config"
	"project-v/pkg/container"
	handleossignal "project-v/pkg/handle-os-signal"
	"project-v/pkg/l"
	"project-v/pkg/l/sentry"
	tracer "project-v/pkg/trace"
)

func main() {
	ll := l.New()
	cfg := config.Load(ll)
	if cfg.SentryConfig.Enabled {
		ll = l.NewWithSentry(&sentry.Configuration{
			DSN: cfg.SentryConfig.DNS,
			Trace: struct{ Disabled bool }{
				Disabled: !cfg.SentryConfig.Trace,
			},
		})
	}
	container.Register(func() l.Logger {
		return ll
	})
	tracer.NewWithOption(
		&tracer.Option{
			TracerName: cfg.Tracing.Name + ":" + cfg.Environment,
			L:          ll,
			TelExporter: &tracer.TelExporter{
				ID:    cfg.Tracing.TeleID,
				Token: cfg.Tracing.TeleToken,
			},
			Enable: true,
		},
	)

	// init os signal handle
	shutdown := handleossignal.New(ll)
	shutdown.HandleDefer(func() {
		ll.Sync()
	})
	container.Register(func() handleossignal.IShutdownHandler {
		return shutdown
	})

	bootstrap(cfg)

	// go registerHttpServer(cfg)

	// handle signal
	if cfg.Environment == "D" {
		shutdown.SetTimeout(1)
	}
	shutdown.Handle()
}
