package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthHandler - ヘルスチェック用の handler
type HealthHandler interface {
	GetHealth(c *gin.Context)
}

type healthHandler struct {
}

// NewHealthCheckHandler - healthCheckHandler の生成
func NewHealthCheckHandler() HealthHandler {
	return &healthHandler{}
}

func (h *healthHandler) GetHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}
