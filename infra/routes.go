package infra

import "github.com/gbzarelli/pgen/controllers"

func ConfigureRoutes(server *GinHttpServer, protocolController *controllers.ProtocolController) {
	mainGroup := server.GetEngine().Group("v1")
	{
		protocolGroup := mainGroup.Group("protocol")
		{
			protocolGroup.POST("/", protocolController.CreateNewProtocol)
		}
	}

}
