package main

import (
	"github.com/gbzarelli/pgen/internal/api"
	"github.com/gbzarelli/pgen/internal/cache"
	"github.com/gbzarelli/pgen/internal/envvar"
	"github.com/gbzarelli/pgen/internal/metrics"
	"github.com/gbzarelli/pgen/internal/service"
)

func main() {
	cacheRedis := cache.NewCacheRedis()

	decimalPlaces := envvar.GetIntegerEnv(envvar.ProtocolDecimalPlacesAfterDateEnv, service.DefaultProtocolDecimalPlacesAfterDate)
	protocolService := service.NewProtocolService(decimalPlaces)
	serviceCache := service.NewProtocolCacheService(protocolService, cacheRedis.GetClient())

	protocolController := api.NewProtocolController(serviceCache)

	ginHTTPServer := api.NewGinHTTPServer()
	prometheus := metrics.NewGinPrometheus(ginHTTPServer.GetEngine())

	api.ConfigureRoutes(
		ginHTTPServer,
		protocolController,
		prometheus.GinPrometheus.Instrument(),
	)

	ginHTTPServer.RunServer()
}
