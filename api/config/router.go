package config

import "github.com/gin-gonic/gin"

// Router - ルーティングの定義
func Router() *gin.Engine {
	registry := NewRegistry()

	// ルーティング
	router := gin.Default()
	router.GET("/", registry.HealthCheckController.GetHealth)

	return router
}
