package systemcontroller

import (
	"github.com/gin-gonic/gin"
	"github.com/yx1126/go-admin/app/service"
	"github.com/yx1126/go-admin/app/vo"
	"github.com/yx1126/go-admin/response"
)

type DictTypeController struct{}

func (*DictTypeController) QueryAllList(c *gin.Context) {
	response.New((&service.SysDictTypeService{}).QueryDictTypeAllList()).Json(c)
}

func (*DictTypeController) Create(c *gin.Context) {
	var dictType vo.CreateDictType
	err := c.ShouldBindJSON(dictType)
	if err != nil {
		response.NewError(err).Json(c)
		return
	}
	response.New(nil, (&service.SysDictTypeService{}).CreateDictType(dictType)).Json(c)
}

func (*DictTypeController) Update(c *gin.Context) {
	var dictType vo.UpdateDictType
	err := c.ShouldBindJSON(dictType)
	if err != nil {
		response.NewError(err).Json(c)
		return
	}
	response.New(nil, (&service.SysDictTypeService{}).UpdateDictType(dictType)).Json(c)
}

func (*DictTypeController) Delete(c *gin.Context) {
	var ids []int
	err := c.ShouldBindJSON(&ids)
	if err != nil {
		response.NewError(err.Error()).Json(c)
		return
	}
	response.New(nil, (&service.SysDictTypeService{}).DeleteDictType(ids)).Json(c)
}
