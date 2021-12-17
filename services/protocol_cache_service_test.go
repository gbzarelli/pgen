package services

import (
	"errors"
	"fmt"
	"github.com/go-redis/redismock/v8"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"strings"
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

func TestNotGenerateNewProtocol(t *testing.T) {
	t.Run("not generate when cache is out", func(t *testing.T) {
		asserts := assert.New(t)
		cacheService, cacheMock, pServiceMock := prepareServicesAndMocks()
		protocol := "fake"
		cacheError := errors.New("out")

		cacheMock.ExpectExists(protocol).SetErr(cacheError)
		for i := 1; i <= MaxRetryGenerateProtocol; i++ {
			cacheMock.ExpectExists(fmt.Sprintf("%s-%d", protocol, i)).SetErr(cacheError)
		}
		pServiceMock.On("NewProtocol").Return(protocol, nil)

		protocol, err := cacheService.NewProtocol()

		asserts.NotNil(err)
		asserts.Equal(cacheError, err)
		pServiceMock.AssertNumberOfCalls(t, "NewProtocol", MaxRetryGenerateProtocol+1)
	})

	t.Run("not generate when always protocols already exists", func(t *testing.T) {
		asserts := assert.New(t)
		cacheService, cacheMock, pServiceMock := prepareServicesAndMocks()
		protocol := "fake"

		setExistsValueInCacheForAllRetries(cacheMock, protocol, 1)
		pServiceMock.On("NewProtocol").Return(protocol, nil)

		protocol, err := cacheService.NewProtocol()

		asserts.NotNil(err)
		asserts.Equal(errProtocolAlreadyExist, err)
		pServiceMock.AssertNumberOfCalls(t, "NewProtocol", MaxRetryGenerateProtocol+1)
	})

	t.Run("not generate when some error in set cache", func(t *testing.T) {
		asserts := assert.New(t)
		cacheService, cacheMock, pServiceMock := prepareServicesAndMocks()
		protocol := "fake"

		setExistsValueInCacheForAllRetries(cacheMock, protocol, 0)
		pServiceMock.On("NewProtocol").Return(protocol, nil)

		protocol, err := cacheService.NewProtocol()

		asserts.NotNil(err)
		asserts.Equal(true, strings.Contains(err.Error(), "call to cmd '[set"))
		pServiceMock.AssertNumberOfCalls(t, "NewProtocol", MaxRetryGenerateProtocol+1)
	})
}

func setExistsValueInCacheForAllRetries(cacheMock redismock.ClientMock, protocol string, value int64) {
	cacheMock.ExpectExists(protocol).SetVal(value)
	for i := 1; i <= MaxRetryGenerateProtocol; i++ {
		cacheMock.ExpectExists(fmt.Sprintf("%s-%d", protocol, i)).SetVal(value)
	}
}

func prepareServicesAndMocks() (ProtocolService, redismock.ClientMock, *MockProtocolService) {
	pServiceMock := new(MockProtocolService)
	db, redisMock := redismock.NewClientMock()
	serviceCache := NewProtocolCacheService(pServiceMock, db)
	return serviceCache, redisMock, pServiceMock
}
