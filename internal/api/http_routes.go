package api

import (
	"github.com/gin-gonic/gin"
)

// ConfigureRoutes Configure HTTP routes in GinHTTPServer
func ConfigureRoutes(server *GinHTTPServer, protocolController *ProtocolController, middleware ...gin.HandlerFunc) {
	server.GetEngine().Use(middleware...)

	mainGroup := server.GetEngine().Group("v1")
	{
		protocolGroup := mainGroup.Group("protocol")
		{
			protocolGroup.POST("/", protocolController.CreateNewProtocol)
		}
	}

}
