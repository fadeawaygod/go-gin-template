package main

import (
	"go-gin-template/config"
	"go-gin-template/handler"
	"log"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	apiv1 := router.Group("/api/v1")
	{
		apiv1.GET("/ping", handler.GetPing)
		apiv1.GET("/tags", handler.GetTags)
		apiv1.POST("/tags", handler.AddTag)
		apiv1.PATCH("/tags/:id", handler.EditTag)
		apiv1.DELETE("/tags/:id", handler.DeleteTag)
	}

	return router
}

func main() {
	gin.SetMode(config.GetEnv("RUN_MODE"))
	r := setupRouter()
	err := r.Run(config.GetEnv("HTTP_PORT"))
	if err != nil {
		log.Fatal("start server failed.")
	}
}
