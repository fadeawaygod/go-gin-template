package handler

import "github.com/gin-gonic/gin"

// GetPing return pong json
func GetPing(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
