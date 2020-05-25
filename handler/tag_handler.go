package handler

import (
	"go-gin-template/model"
	"go-gin-template/util"

	"github.com/gin-gonic/gin"
)

func GetTags(c *gin.Context) {

	skip, err := util.ReadIntParameter(c, "skip", false, 0)
	if err != nil {
		util.WriteError(c, err)
		return
	}
	limit, err := util.ReadIntParameter(c, "limit", false, 0)
	if err != nil {
		util.WriteError(c, err)
		return
	}
	optioanlParameters := []string{"name", "state"}
	maps := util.ParameterToMap(c, optioanlParameters)
	data := make(map[string]interface{})

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
