package metrics

import (
	"github.com/Depado/ginprom"
	"github.com/gin-gonic/gin"
)

type Prometheus struct {
	GinPrometheus *ginprom.Prometheus
}

// NewGinPrometheus Crate a new instance of Prometheus configured to gin
func NewGinPrometheus(gin *gin.Engine) *Prometheus {
	return &Prometheus{GinPrometheus: ginprom.New(
		ginprom.Engine(gin),
		ginprom.Subsystem("gin"),
		ginprom.Path("/metrics"),
	)}
}
