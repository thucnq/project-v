package redis

import (
	"github.com/go-redis/redis/v8"
)

// NewRedisStore ...
func NewRedisStore(cfg *Config) (*redis.Client, error) {
	opts1, opts2, err := cfg.Parse()
	if err != nil {
		return nil, err
	}

	var client *redis.Client
	if opts1 != nil {
		client = redis.NewClient(opts1)
	} else {
		client = redis.NewFailoverClient(opts2)
	}
	// ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	// defer cancel()
	// client.Ping(ctx)
	// if err = client.Ping(ctx).Err(); err != nil {
	// 	return nil, err
	// }
	return client, nil
}

// MustNewRedisStore ...
func MustNewRedisStore(cfg *Config) *redis.Client {
	s, err := NewRedisStore(cfg)
	if err != nil {
		panic(err)
	}
	return s
}
