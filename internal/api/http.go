package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

// GinHTTPServer struct to manage GinHTTPServer
type GinHTTPServer struct {
	port   uint16
	engine *gin.Engine
}

// NewGinHTTPServer Create a new instance of GinHTTPServer
func NewGinHTTPServer() *GinHTTPServer {
	return &GinHTTPServer{
		port:   5000,
		engine: gin.Default(),
	}
}

// RunServer Just Run Server
func (httpServer *GinHTTPServer) RunServer() {
	serverAddress := ":" + strconv.Itoa(int(httpServer.port))
	log.Fatal(httpServer.engine.Run(serverAddress))
}

// GetEngine Get the gin.Engine
func (httpServer *GinHTTPServer) GetEngine() *gin.Engine {
	return httpServer.engine
}
