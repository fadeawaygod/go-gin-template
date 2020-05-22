package main

import (
	"go-gin-template/config"
	"go-gin-template/handler"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", handler.GetPing)

	return r
}

func main() {
	r := setupRouter()
	r.Run(config.GetEnv("GIN_HOST")) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
