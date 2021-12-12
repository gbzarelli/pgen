package main

import (
	"github.com/gbzarelli/pgen/controllers"
	"github.com/gbzarelli/pgen/infra"
	"github.com/gbzarelli/pgen/services"
)

const ProtocolDecimalPlacesAfterDateEnv = "PROTOCOL_DECIMAL_PLACES_AFTER_DATE"

func main() {
	cache := infra.NewCacheRedis()

	decimalPlaces := infra.GetIntegerEnv(ProtocolDecimalPlacesAfterDateEnv, services.DefaultProtocolDecimalPlacesAfterDate)
	service := services.NewProtocolService(decimalPlaces)
	serviceCache := services.NewProtocolCacheService(service, cache.GetClient())

	protocolController := controllers.NewProtocolController(serviceCache)

	ginHttpServer := infra.NewGinHttpServer()
	infra.ConfigureRoutes(ginHttpServer, protocolController)
	ginHttpServer.RunServer()
}
