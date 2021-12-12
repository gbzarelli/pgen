package main

import (
	"github.com/gbzarelli/pgen/controllers"
	"github.com/gbzarelli/pgen/infra"
	"github.com/gbzarelli/pgen/services"
)

func main() {
	cache := infra.NewCacheRedis()

	service := services.NewProtocolService(services.DefaultProtocolDecimalPlacesAfterDate)
	serviceCache := services.NewProtocolCacheService(service, cache.GetClient())

	protocolController := controllers.NewProtocolController(serviceCache)

	ginHttpServer := infra.NewGinHttpServer()
	infra.ConfigureRoutes(ginHttpServer, protocolController)
	ginHttpServer.RunServer()
}
