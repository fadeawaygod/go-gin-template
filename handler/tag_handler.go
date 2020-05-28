package handler

import (
	"go-gin-template/exception"
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

	data["lists"], err = model.GetTags(skip, limit, maps)
	if err != nil {
		util.WriteError(c, err)
		return
	}
	data["total"], err = model.GetTagTotal(maps)
	if err != nil {
		util.WriteError(c, err)
		return
	}
	util.WriteSuccess(c, data)
}

func AddTag(c *gin.Context) {
	var body model.RawTag
	ginErr := c.BindJSON(&body)
	if ginErr != nil {
		util.WriteError(c, &exception.ParseJsonBodyError)
		return
	}
	err := model.AddTag(body)
	if err != nil {
		util.WriteError(c, err)
		return
	}
	util.WriteSuccess(c, nil)
}

func EditTag(c *gin.Context) {
	var body model.RawTag
	tagId, err := util.ReadIntURLParameter(c, "id")
	if err != nil {
		util.WriteError(c, err)
		return
	}

	ginErr := c.BindJSON(&body)
	if ginErr != nil {
		util.WriteError(c, &exception.ParseJsonBodyError)
		return
	}

	err = model.EditTag(tagId, body)
	if err != nil {
		util.WriteError(c, err)
		return
	}
	util.WriteSuccess(c, nil)
}

func DeleteTag(c *gin.Context) {
}
