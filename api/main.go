package main

import (
	"log"
	"net/http"

	"github.com/16francs/task-management/api/config"
)

func main() {
	// 起動設定
	router := config.Router()
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("alert: %v", err)
	}
}
