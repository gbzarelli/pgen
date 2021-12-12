package infra

import (
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

type GinHttpServer struct {
	port   uint16
	engine *gin.Engine
}

func NewGinHttpServer() *GinHttpServer {
	return &GinHttpServer{
		port:   5000,
		engine: gin.Default(),
	}
}

func (httpServer *GinHttpServer) RunServer() {
	serverAddress := ":" + strconv.Itoa(int(httpServer.port))
	log.Fatal(httpServer.engine.Run(serverAddress))
}

func (httpServer *GinHttpServer) GetEngine() *gin.Engine {
	return httpServer.engine
}
