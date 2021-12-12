package controllers

import (
	"github.com/gbzarelli/pgen/services"
	"github.com/gin-gonic/gin"
)

type ProtocolController struct {
	service services.ProtocolService
}

func NewProtocolController(service services.ProtocolService) *ProtocolController {
	return &ProtocolController{service: service}
}

func (protocolController ProtocolController) CreateNewProtocol(context *gin.Context) {
	protocol, err := protocolController.service.NewProtocol()
	if err == nil {
		context.JSON(201, gin.H{
			"protocol": protocol,
		})
	} else {
		context.JSON(500, gin.H{
			"error": err.Error(),
		})
	}
}
