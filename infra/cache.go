package infra

import (
	"github.com/go-redis/redis/v8"
)

// CacheRedis struct to manage the cache
type CacheRedis struct {
	opts   *redis.Options
	client *redis.Client
}

// NewCacheRedis Create a new instance of Redis using redis.NewClient
func NewCacheRedis() *CacheRedis {
	opts := &redis.Options{
		Addr:     GetStringEnv(RedisAddressEnv, "localhost:6379"),
		Password: "", // no password set
		DB:       0,  // use default DB
	}
	return &CacheRedis{
		opts:   opts,
		client: redis.NewClient(opts),
	}
}

// GetClient Return the redis.Client
func (cr *CacheRedis) GetClient() *redis.Client {
	return cr.client
}
