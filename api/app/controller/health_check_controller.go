package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthCheckController - ヘルスチェック用 Controller
type HealthCheckController interface {
	GetHealth(c *gin.Context)
}

type healthCheckController struct {
}

// NewHealthCheckController - healthCheckController の生成
func NewHealthCheckController() HealthCheckController {
	return &healthCheckController{}
}

func (h *healthCheckController) GetHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}
