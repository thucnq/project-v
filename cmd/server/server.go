package main

import (
	"project-v/config"
	api "project-v/internal/app-api"
	"project-v/internal/app-api/handler/health"
	"project-v/pkg/container"
	handleossignal "project-v/pkg/handle-os-signal"
	"project-v/pkg/l"
)

func registerHttpServer(cfg *config.Config) {
	var ll = container.ResolverMust[l.Logger]()
	var healthHandler = container.ResolverMust[health.Controller]()
	var shutdown = container.ResolverMust[handleossignal.IShutdownHandler]()

	healthHandler.SetReady(true)
	shutdown.HandleDefer(
		func() {
			healthHandler.SetReady(false)
		},
	)

	gw := api.New(cfg.Environment)
	gw.Middleware(ll)
	gw.InitHealth(healthHandler)
	gw.InitLogHandler()
	gw.InitMetrics()
	ll.Info(
		"HTTP server start listening",
		l.Any("HTTP address", cfg.GetHTTPAddress()),
	)
	err := gw.Listen(cfg.GetHTTPAddress())
	if err != nil {
		ll.Fatal(
			"error listening to address",
			l.Any("address", cfg.GetHTTPAddress()), l.Error(err),
		)
		return
	}
}
