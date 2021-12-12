package infra

import (
	"github.com/go-redis/redis/v8"
)

type CacheRedis struct {
	opts   *redis.Options
	client *redis.Client
}

func NewCacheRedis() *CacheRedis {
	opts := &redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	}
	return &CacheRedis{
		opts:   opts,
		client: redis.NewClient(opts),
	}
}

func (cr *CacheRedis) GetClient() *redis.Client {
	return cr.client
}
