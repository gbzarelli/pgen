package api

// ConfigureRoutes Configure HTTP routes in GinHTTPServer
func ConfigureRoutes(server *GinHTTPServer, protocolController *ProtocolController) {
	mainGroup := server.GetEngine().Group("v1")
	{
		protocolGroup := mainGroup.Group("protocol")
		{
			protocolGroup.POST("/", protocolController.CreateNewProtocol)
		}
	}

}
