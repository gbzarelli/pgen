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
	return p.newProtocol(MaxRetryGenerateProtocol)
}

func (p *ProtocolCacheService) newProtocol(retryCount int8) (string, error) {
	protocol, _ := p.targetService.NewProtocol()
	exists, err := p.cache.Exists(ctx, protocol).Result()

	if exists > 0 {
		err = protocolAlreadyExistError
	}

	if err == nil {
		var result string
		result, err = p.cache.Set(ctx, protocol, true, Day).Result()
		if result == CacheSuccessResult {
			return protocol, nil
		}
	}

	if retryCount <= 0 {
		return "", err
	}

	return p.newProtocol(retryCount - 1)
}
