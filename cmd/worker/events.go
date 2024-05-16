package main

import (
	_ "embed"
	"encoding/json"

	"project-v/config"
	"project-v/pkg/worker"
)

//go:embed events.json
var eventBytes []byte

func loadEvents(cfg *config.Config) error {
	var events []*worker.ConsumerHandlerConfig
	if err := json.Unmarshal(eventBytes, &events); err != nil {
		return err
	}
	cfg.RabbitmqConfig.Events = events
	return nil
}
