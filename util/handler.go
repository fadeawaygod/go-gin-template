package util

import (
	"go-gin-template/error"
	"net/http"

	"github.com/gin-gonic/gin"
)

func WriteSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data":   data,
	})
}
func WriteError(c *gin.Context, err *(error.Error)) {
	c.JSON(err.HttpStatusCode, gin.H{
		"status":  "error",
		"code":    err.Code,
		"message": err.Message,
	})
}
