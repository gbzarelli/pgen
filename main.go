package main

import (
	"github.com/gbzarelli/pgen/controllers"
	"github.com/gbzarelli/pgen/infra"
	"github.com/gbzarelli/pgen/services"
)

func main() {
	cache := infra.NewCacheRedis()

	decimalPlaces := infra.GetIntegerEnv(infra.ProtocolDecimalPlacesAfterDateEnv, services.DefaultProtocolDecimalPlacesAfterDate)
	service := services.NewProtocolService(decimalPlaces)
	serviceCache := services.NewProtocolCacheService(service, cache.GetClient())

	protocolController := controllers.NewProtocolController(serviceCache)

	ginHttpServer := infra.NewGinHTTPServer()
	infra.ConfigureRoutes(ginHttpServer, protocolController)
	ginHttpServer.RunServer()
}
