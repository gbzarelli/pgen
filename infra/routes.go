package infra

import "github.com/gbzarelli/pgen/controllers"

// ConfigureRoutes Configure HTTP routes in GinHTTPServer
func ConfigureRoutes(server *GinHTTPServer, protocolController *controllers.ProtocolController) {
	mainGroup := server.GetEngine().Group("v1")
	{
		protocolGroup := mainGroup.Group("protocol")
		{
			protocolGroup.POST("/", protocolController.CreateNewProtocol)
		}
	}

}
