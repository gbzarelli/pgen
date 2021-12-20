package api

import (
	"github.com/gbzarelli/pgen/internal/service"
	"github.com/gin-gonic/gin"
)

// ProtocolController struct to manage the protocol controller
type ProtocolController struct {
	service service.ProtocolService
}

// NewProtocolController Create a new instance of ProtocolController
func NewProtocolController(service service.ProtocolService) *ProtocolController {
	return &ProtocolController{service: service}
}

// CreateNewProtocol Manage the gin.Context to create a new protocol
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
