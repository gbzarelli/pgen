package services

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"time"
)

const Day = time.Hour * 24
const MaxRetryGenerateProtocol = 10
const CacheSuccessResult = "OK"

var (
	ctx                       = context.TODO()
	protocolAlreadyExistError = errors.New("impossible to generate a non-existent protocol")
)

type ProtocolCacheService struct {
	targetService ProtocolService
	cache         *redis.Client
}

func NewProtocolCacheService(service ProtocolService, cache *redis.Client) ProtocolService {
	return &ProtocolCacheService{
		targetService: service,
		cache:         cache,
	}
}

func (p *ProtocolCacheService) NewProtocol() (string, error) {
	var methodError error

	for i := 0; i < MaxRetryGenerateProtocol; i++ {
		methodError = nil
		protocol, _ := p.targetService.NewProtocol()

		exists, cacheExistsErr := p.cache.Exists(ctx, protocol).Result()
		if exists > 0 {
			methodError = protocolAlreadyExistError
		}
		if cacheExistsErr != nil {
			methodError = cacheExistsErr
			continue
		}

		result, cacheSetErr := p.cache.Set(ctx, protocol, true, Day).Result()
		if result == CacheSuccessResult {
			return protocol, nil
		}
		methodError = cacheSetErr
	}

	return "", methodError
}
