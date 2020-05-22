package handler

import (
	"go-gin-template/error"
	"go-gin-template/model"
	"go-gin-template/util"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetTags(c *gin.Context) {
	name := c.Query("name")
	skipRaw := c.Query("skip")
	limitRaw := c.Query("limit")
	skip, err := strconv.Atoi(skipRaw)
	if err != nil {
		err := error.FormatError(error.MissongRequiredParameterError, "skip")
		util.WriteError(c, &err)
		return
	}
	limit, err := strconv.Atoi(limitRaw)
	if err != nil {
	}

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}
	if arg := c.Query("state"); arg != "" {
		maps["state"] = -1
	}

	data["lists"] = model.GetTags(skip, limit, maps)
	data["total"] = model.GetTagTotal(maps)
	util.WriteSuccess(c, data)
}

func AddTag(c *gin.Context) {
}

func EditTag(c *gin.Context) {
}

func DeleteTag(c *gin.Context) {
}
