package config

import "github.com/16francs/task-management/api/app/controller"

// Registry - 簡易の DI コンテナ
// router の実装が依存する controller を定義
type Registry struct {
	controller.HealthCheckController
}

// NewRegistry - registry の生成
func NewRegistry() Registry {
	healthCheckController := controller.NewHealthCheckController()

	return Registry{
		healthCheckController,
	}
}
