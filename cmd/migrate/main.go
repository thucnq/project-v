package main

import (
	"project-v/config"
	"project-v/database"
	"project-v/pkg/l"
)

func main() {
	ll := l.New()
	cfg := config.Load(ll)

	err := database.RunMysqlMigrate(&cfg.MySql)
	if err != nil {
		ll.Error("Failed to migrate", l.Error(err))
		return
	}

	ll.Info("Migrate success")
}
