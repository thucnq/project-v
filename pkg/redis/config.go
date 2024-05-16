package redis

import (
	"errors"
	"time"

	"github.com/go-redis/redis/v8"
)

// Config represents generic config for both standalone server and sentinel
type Config struct {
	UseSentinel bool     `json:"use_sentinel" mapstructure:"use_sentinel"`
	MasterName  string   `json:"master" mapstructure:"master"`
	Addrs       []string `json:"address" mapstructure:"address"`

	Password string `json:"password" mapstructure:"password"`
	Database int    `json:"database" mapstructure:"database"`

	MaxRetries      int           `json:"max_retries" mapstructure:"max_retries"`
	MinRetryBackoff time.Duration `json:"min_retry_backoff" mapstructure:"min_retry_backoff"`
	MaxRetryBackoff time.Duration `json:"max_retry_backoff" mapstructure:"max_retry_backoff"`

	DialTimeout  time.Duration `json:"dial_timeout" mapstructure:"dial_timeout"`
	ReadTimeout  time.Duration `json:"read_timeout" mapstructure:"read_timeout"`
	WriteTimeout time.Duration `json:"write_timeout" mapstructure:"write_timeout"`

	PoolSize           int           `json:"pool_size" mapstructure:"pool_size"`
	MinIdleConns       int           `json:"min_idle_conns" mapstructure:"min_idle_conns"`
	MaxConnAge         time.Duration `json:"max_conn_age" mapstructure:"max_conn_age"`
	PoolTimeout        time.Duration `json:"pool_timeout" mapstructure:"pool_timeout"`
	IdleTimeout        time.Duration `json:"idle_timeout" mapstructure:"idle_timeout"`
	IdleCheckFrequency time.Duration `json:"idle_check_frequency" mapstructure:"idle_check_frequency"`
}

// Parse handles config for both standalone server and sentinel
func (c *Config) Parse() (*redis.Options, *redis.FailoverOptions, error) {
	if c.UseSentinel {
		if c.MasterName == "" {
			return nil, nil, errors.New("pkg/redis: invalid config, must provide master name when using sentinel")
		}
		if len(c.Addrs) == 0 {
			return nil, nil, errors.New("pkg/redis: invalid config, must provide address when using sentinel")
		}

	} else {
		if len(c.Addrs) != 1 {
			return nil, nil, errors.New("pkg/redis: invalid config, must provide exactly 1 address when using standalone")
		}
	}

	if !c.UseSentinel {
		opts, err := redis.ParseURL(c.Addrs[0])
		if err != nil {
			return nil, nil, err
		}
		if len(c.Password) > 0 && len(opts.Password) == 0 {
			opts.Password = c.Password
		}
		if opts.DB == 0 && c.Database > 0 {
			opts.DB = c.Database
		}

		opts.PoolSize = 1000

		return opts, nil, err
	}

	opts := &redis.FailoverOptions{
		MasterName:         c.MasterName,
		SentinelAddrs:      c.Addrs,
		OnConnect:          nil,
		Password:           c.Password,
		DB:                 c.Database,
		MaxRetries:         c.MaxRetries,
		MinRetryBackoff:    c.MinRetryBackoff,
		MaxRetryBackoff:    c.MaxRetryBackoff,
		DialTimeout:        c.DialTimeout,
		ReadTimeout:        c.ReadTimeout,
		WriteTimeout:       c.WriteTimeout,
		PoolSize:           c.PoolSize,
		MinIdleConns:       c.MinIdleConns,
		MaxConnAge:         c.MaxConnAge,
		PoolTimeout:        c.PoolTimeout,
		IdleTimeout:        c.IdleTimeout,
		IdleCheckFrequency: c.IdleCheckFrequency,
		TLSConfig:          nil,
	}
	return nil, opts, nil
}

// DefaultConfig returns default config for testing
func DefaultConfig() *Config {
	return &Config{
		UseSentinel: false,
		Addrs:       []string{"redis://redis:6379"},
	}
}
