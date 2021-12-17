package services

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"time"
)

// Day represents a Day in nano
const Day = time.Hour * 24

// MaxRetryGenerateProtocol value to max retry in generate new protocol when some error happens
const MaxRetryGenerateProtocol = 10

// CacheSuccessResult value to represent the cache success response
const CacheSuccessResult = "OK"

var (
	ctx                     = context.TODO()
	errProtocolAlreadyExist = errors.New("impossible to generate a non-existent protocol")
)

// ProtocolCacheService struct to manage the protocol cache service
type ProtocolCacheService struct {
	targetService ProtocolService
	cache         *redis.Client
}

// NewProtocolCacheService Create a new instance of ProtocolCacheService
func NewProtocolCacheService(service ProtocolService, cache *redis.Client) ProtocolService {
	return &ProtocolCacheService{
		targetService: service,
		cache:         cache,
	}
}

// NewProtocol method to generate a new protocol
func (p *ProtocolCacheService) NewProtocol() (string, error) {
	return p.newProtocol(MaxRetryGenerateProtocol)
}

func (p *ProtocolCacheService) newProtocol(retryCount int8) (string, error) {
	protocol, _ := p.targetService.NewProtocol()
	exists, err := p.cache.Exists(ctx, protocol).Result()

	if exists > 0 {
		err = errProtocolAlreadyExist
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
