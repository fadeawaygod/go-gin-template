package util

import (
	"go-gin-template/exception"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func WriteSuccess(c *gin.Context, data interface{}) {
	var body = make(gin.H)
	body["status"] = "ok"
	if data != nil {
		body["data"] = data
	}

	c.JSON(http.StatusOK, body)
}
func WriteError(c *gin.Context, err *(exception.Exception)) {
	c.JSON(err.HttpStatusCode, gin.H{
		"status":  "error",
		"code":    err.Code,
		"message": err.Message,
	})
}

// ReadIntParameter read an url parameter and cast into integer
func ReadIntParameter(c *gin.Context, name string, isRequired bool, defaultValue int) (int, *exception.Exception) {
	rawString := c.Query(name)
	value, err := strconv.Atoi(rawString)
	if err != nil {
		if isRequired {
			formatedException := exception.FormatException(exception.MissongRequiredParameterError, name)
			return 0, &formatedException
		}
		return defaultValue, nil
	}
	return value, nil
}

// ReadIntParameter read an url parameter and cast into integer
func ReadIntURLParameter(c *gin.Context, name string) (int, *exception.Exception) {
	rawString := c.Param(name)
	value, err := strconv.Atoi(rawString)
	if err != nil {
		formatedException := exception.FormatException(exception.MissongURLParameterError, name)
		return 0, &formatedException
	}
	return value, nil
}

func ParameterToMap(c *gin.Context, paras []string) map[string]interface{} {
	mapResult := make(map[string]interface{})
	for _, para := range paras {
		if arg := c.Query(para); arg != "" {
			mapResult[para] = arg
		}
	}
	return mapResult
}
