// YOU CAN EDIT YOUR CUSTOM CONFIG HERE

package config

import (
	"fmt"
	"strconv"

	"project-v/pkg/mysql"
	"project-v/pkg/redis"
	"project-v/pkg/worker"
)

// Config ...
type Config struct {
	Base         `mapstructure:",squash"`
	SentryConfig SentryConfig `json:"sentry_config" mapstructure:"sentry_config"`

	MySql          mysql.Config   `json:"mysql" mapstructure:"mysql"`
	UserRedis      redis.Config   `json:"user_redis" mapstructure:"user_redis"`
	RabbitmqConfig RabbitmqConfig `json:"rabbitmq" mapstructure:"rabbitmq"`
	Redis          Redis          `json:"redis" mapstructure:"redis"`

	CDNMappingFilePath string                `json:"cdn_mapping_filepath" mapstructure:"cdn_mapping_filepath"`
	ScheduleConfig     worker.ScheduleConfig `json:"schedule" mapstructure:"schedule"`
}

// GetHTTPAddress ...
func (c *Config) GetHTTPAddress() string {
	if _, err := strconv.Atoi(c.HTTPAddress); err == nil {
		return fmt.Sprintf(":%v", c.HTTPAddress)
	}
	return c.HTTPAddress
}

// SentryConfig ...
type SentryConfig struct {
	Enabled bool   `json:"enabled" mapstructure:"enabled"`
	DNS     string `json:"dns" mapstructure:"dns"`
	Trace   bool   `json:"trace" mapstructure:"trace"`
}

// RabbitmqConfig ...
type RabbitmqConfig struct {
	Enabled           bool   `json:"enabled" mapstructure:"enabled"`
	URI               string `json:"uri" mapstructure:"uri"`
	Retries           int    `json:"retries" mapstructure:"retries"`
	InternalQueueSize int    `json:"internal_queue_size" mapstructure:"internal_queue_size"`
	MaxThread         int    `json:"max_thread" mapstructure:"max_thread"`
	// Exchange the exchange name for publisher
	Exchange string `json:"exchange" mapstructure:"exchange"`
	// ExchangeType the exchange type for publisher
	ExchangeType string `json:"exchange_type" mapstructure:"exchange_type"`
	// Events the list event to listen
	Events []*worker.ConsumerHandlerConfig `json:"events" mapstructure:"events"`
}

type Redis struct {
	Address   string `json:"address" mapstructure:"address"`
	Namespace string `json:"namespace" mapstructure:"namespace"`
	Database  int    `json:"database" mapstructure:"database"`
	Password  string `json:"password" mapstructure:"password"`
}
