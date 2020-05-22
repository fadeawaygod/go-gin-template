package main

import (
	"go-gin-template/config"
	"go-gin-template/handler"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/ping", handler.GetPing)

	return router
}

func main() {
	gin.SetMode(config.GetEnv("RUN_MODE"))
	r := setupRouter()
	r.Run(config.GetEnv("HTTP_PORT"))
}
