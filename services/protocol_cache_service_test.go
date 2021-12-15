package services

import (
	"fmt"
	"github.com/go-redis/redismock/v8"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockProtocolService struct {
	mock.Mock
}

func (mock *MockProtocolService) NewProtocol() (string, error) {
	args := mock.Called()
	result := args.Get(0)

	callsCount := len(mock.Calls)
	if callsCount == 1 {
		return result.(string), args.Error(1)
	}

	resultAfterFirstCall := fmt.Sprintf("%s-%d", result, callsCount-1)
	return resultAfterFirstCall, args.Error(1)
}

func TestGenerateNewProtocolWithoutExistsInCache(t *testing.T) {
	asserts := assert.New(t)
	cacheService, cacheMock, pServiceMock := prepareServicesAndMocks()
	expectedProtocol := "123456789"

	cacheMock.ExpectExists(expectedProtocol).SetVal(0)
	cacheMock.ExpectSet(expectedProtocol, true, Day).SetVal(CacheSuccessResult)
	pServiceMock.On("NewProtocol").Return(expectedProtocol, nil)

	protocol, err := cacheService.NewProtocol()

	asserts.Equal(expectedProtocol, protocol)
	asserts.Nil(err)
	pServiceMock.AssertNumberOfCalls(t, "NewProtocol", 1)
}

func TestGenerateNewProtocolWithExistsInCache(t *testing.T) {
	t.Run("with one retry", func(t *testing.T) {
		asserts := assert.New(t)
		cacheService, cacheMock, pServiceMock := prepareServicesAndMocks()
		existedProtocol := "123456789"
		expectedProtocol := "123456789-1"

		cacheMock.ExpectExists(existedProtocol).SetVal(1)
		cacheMock.ExpectExists(expectedProtocol).SetVal(0)
		cacheMock.ExpectSet(expectedProtocol, true, Day).SetVal(CacheSuccessResult)
		pServiceMock.On("NewProtocol").Return(existedProtocol, nil)

		protocol, err := cacheService.NewProtocol()

		asserts.Equal(expectedProtocol, protocol)
		asserts.Nil(err)
		pServiceMock.AssertNumberOfCalls(t, "NewProtocol", 2)
	})
	t.Run("with four retry", func(t *testing.T) {
		asserts := assert.New(t)
		cacheService, cacheMock, pServiceMock := prepareServicesAndMocks()
		existedProtocol := "123456789"
		expectedProtocol := "123456789-4"

		cacheMock.ExpectExists(existedProtocol).SetVal(1)
		for i := 1; i < 4; i++ {
			cacheMock.ExpectExists(fmt.Sprintf("%s-%d", existedProtocol, i)).SetVal(1)
		}
		cacheMock.ExpectExists(expectedProtocol).SetVal(0)
		cacheMock.ExpectSet(expectedProtocol, true, Day).SetVal(CacheSuccessResult)
		pServiceMock.On("NewProtocol").Return(existedProtocol, nil)

		protocol, err := cacheService.NewProtocol()

		asserts.Equal(expectedProtocol, protocol)
		asserts.Nil(err)
		pServiceMock.AssertNumberOfCalls(t, "NewProtocol", 5)
	})
}

func prepareServicesAndMocks() (ProtocolService, redismock.ClientMock, *MockProtocolService) {
	pServiceMock := new(MockProtocolService)
	db, redisMock := redismock.NewClientMock()
	serviceCache := NewProtocolCacheService(pServiceMock, db)
	return serviceCache, redisMock, pServiceMock
}
