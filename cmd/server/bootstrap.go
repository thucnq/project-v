package main

import (
	"context"
	"project-v/config"
	"project-v/internal/app-api/handler/health"
	"project-v/pkg/container"
	handleossignal "project-v/pkg/handle-os-signal"
	"project-v/pkg/l"
)

func bootstrap(cfg *config.Config) {
	var _ = container.ResolverMust[l.Logger]()
	var shutdown = container.ResolverMust[handleossignal.IShutdownHandler]()

	loadI18n()

	container.Register(
		func() health.Controller {
			return health.New()
		},
	)

	_, cancel := context.WithCancel(context.Background())
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
	// endregion

	// region register factory
	// endregion

	// region register publisher
	// endregion

	// region register service to container

	// endregion
}
