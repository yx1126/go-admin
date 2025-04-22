package systemcontroller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	service "github.com/yx1126/go-admin/app/service/system"
	"github.com/yx1126/go-admin/app/vo"
	"github.com/yx1126/go-admin/response"
)

type DictController struct{}

func (*DictController) QueryAllList(c *gin.Context) {
	response.New((&service.SysDictTypeService{}).QueryDictTypeAllList()).Json(c)
}

func (*DictController) Create(c *gin.Context) {
	var dictType vo.CreateDictType
	err := c.ShouldBindJSON(&dictType)
	if err != nil {
		response.NewError(err).Json(c)
		return
	}
	if ok := (&service.SysDictTypeService{}).DictTypeHasSameType(dictType.Type, nil); ok {
		response.NewError(nil).SetMsg("字典类型已存在").Json(c)
		return
	}
	response.New(nil, (&service.SysDictTypeService{}).CreateDictType(dictType)).Json(c)
}

func (*DictController) Update(c *gin.Context) {
	var dictType vo.UpdateDictType
	err := c.ShouldBindJSON(&dictType)
	if err != nil {
		response.NewError(err).Json(c)
		return
	}
	dictTypeId := int(dictType.Id)
	if ok := (&service.SysDictTypeService{}).DictTypeHasSameType(dictType.Type, &dictTypeId); ok {
		response.NewError(nil).SetMsg("字典类型已存在").Json(c)
		return
	}
	response.New(nil, (&service.SysDictTypeService{}).UpdateDictType(dictType)).Json(c)
}

func (*DictController) Delete(c *gin.Context) {
	var ids []int
	err := c.ShouldBindJSON(&ids)
	if err != nil {
		response.NewError(err.Error()).Json(c)
		return
	}
	response.New(nil, (&service.SysDictTypeService{}).DeleteDictType(ids)).Json(c)
}

func (*DictController) QueryDictDataList(c *gin.Context) {
	var dictId *int
	if val := c.Query("id"); val != "" {
		if val2, ok := strconv.Atoi(val); ok == nil {
			dictId = &val2
		}
	}
	response.New((&service.SysDictDataService{}).QueryDictDataAllList(dictId)).Json(c)
}

func (*DictController) CreateData(c *gin.Context) {
	var dictData vo.CreateDictData
	err := c.ShouldBindJSON(&dictData)
	if err != nil {
		response.NewError(err).Json(c)
		return
	}
	if ok := (&service.SysDictDataService{}).DictDataHasSameNameValue(dictData.Label, dictData.Value, nil); ok {
		response.NewError(nil).SetMsg("字典名称或值已存在").Json(c)
		return
	}
	response.New(nil, (&service.SysDictDataService{}).CreateDictData(dictData)).Json(c)
}

func (*DictController) UpdateData(c *gin.Context) {
	var dictData vo.UpdateDictData
	err := c.ShouldBindJSON(&dictData)
	if err != nil {
		response.NewError(err).Json(c)
		return
	}
	dictDataId := int(dictData.Id)
	if ok := (&service.SysDictDataService{}).DictDataHasSameNameValue(dictData.Label, dictData.Value, &dictDataId); ok {
		response.NewError(nil).SetMsg("字典名称或值已存在").Json(c)
		return
	}
	response.New(nil, (&service.SysDictDataService{}).UpdateDictData(dictData)).Json(c)
}

func (*DictController) DeleteData(c *gin.Context) {
	var ids []int
	err := c.ShouldBindJSON(&ids)
	if err != nil {
		response.NewError(err.Error()).Json(c)
		return
	}
	response.New(nil, (&service.SysDictDataService{}).DeleteDictData(ids)).Json(c)
}
