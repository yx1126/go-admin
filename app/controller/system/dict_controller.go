package systemcontroller

import (
	"github.com/gin-gonic/gin"
	systemservice "github.com/yx1126/go-admin/app/service/system"
	"github.com/yx1126/go-admin/app/vo"
	bind "github.com/yx1126/go-admin/common/should_bind"
	"github.com/yx1126/go-admin/response"
)

type DictController struct{}

// 字典类型查询
func (*DictController) QueryAllList(c *gin.Context) {
	response.New((&systemservice.SysDictTypeService{}).QueryDictTypeAllList()).Json(c)
}

// 字典类型新增
func (*DictController) Create(c *gin.Context) {
	var dictType vo.CreateDictType
	err := c.ShouldBindJSON(&dictType)
	if err != nil {
		response.NewError(err).Json(c)
		return
	}
	if ok := (&systemservice.SysDictTypeService{}).IsHasSameType(dictType.Type, nil); ok {
		response.NewError(nil).SetMsg("字典类型已存在").Json(c)
		return
	}
	response.New(nil, (&systemservice.SysDictTypeService{}).CreateDictType(dictType)).Json(c)
}

// 字典类型更新
func (*DictController) Update(c *gin.Context) {
	var dictType vo.UpdateDictType
	err := c.ShouldBindJSON(&dictType)
	if err != nil {
		response.NewError(err).Json(c)
		return
	}
	dictTypeId := int(dictType.Id)
	if ok := (&systemservice.SysDictTypeService{}).IsHasSameType(dictType.Type, &dictTypeId); ok {
		response.NewError(nil).SetMsg("字典类型已存在").Json(c)
		return
	}
	response.New(nil, (&systemservice.SysDictTypeService{}).UpdateDictType(dictType)).Json(c)
}

// 字典类型删除
func (*DictController) Delete(c *gin.Context) {
	var ids []int
	err := c.ShouldBindJSON(&ids)
	if err != nil {
		response.NewError(err).Json(c)
		return
	}
	response.New(nil, (&systemservice.SysDictTypeService{}).DeleteDictType(ids)).Json(c)
}

// 根据字典类型id查询字典数据
func (*DictController) QueryDictDataList(c *gin.Context) {
	var params vo.DictPagingParam
	if err := bind.BindPaging(c, &params); err != nil {
		response.NewError(err).Json(c)
		return
	}
	data, err := (&systemservice.SysDictDataService{}).QueryDictDataList(params)
	paging := response.Paging{
		List:  data.Data,
		Page:  params.Page,
		Size:  params.Size,
		Total: data.Count,
	}
	if err != nil {
		response.NewError(err).SetPaging(paging).Json(c)
		return
	}
	response.NewSuccess(paging).Json(c)
}

// 根据字典类型查询字典数据
func (*DictController) QueryDictDataListByType(c *gin.Context) {
	response.New((&systemservice.SysDictDataService{}).QueryDictDataListByType(c.Param("dictType"))).Json(c)
}

// 字典数据新增
func (*DictController) CreateData(c *gin.Context) {
	var dictData vo.CreateDictData
	err := c.ShouldBindJSON(&dictData)
	if err != nil {
		response.NewError(err).Json(c)
		return
	}
	if ok := (&systemservice.SysDictDataService{}).IsHasSameName(dictData.Label, dictData.DictId, nil); ok {
		response.NewError(nil).SetMsg("字典名称已存在").Json(c)
		return
	}
	if ok := (&systemservice.SysDictDataService{}).IsHasSameValue(dictData.Value, dictData.DictId, nil); ok {
		response.NewError(nil).SetMsg("字典值已存在").Json(c)
		return
	}
	response.New(nil, (&systemservice.SysDictDataService{}).CreateDictData(dictData)).Json(c)
}

// 字典数据更新
func (*DictController) UpdateData(c *gin.Context) {
	var dictData vo.UpdateDictData
	err := c.ShouldBindJSON(&dictData)
	if err != nil {
		response.NewError(err).Json(c)
		return
	}
	dictDataId := int(dictData.Id)
	if ok := (&systemservice.SysDictDataService{}).IsHasSameName(dictData.Label, dictData.DictId, &dictDataId); ok {
		response.NewError(nil).SetMsg("字典名称已存在").Json(c)
		return
	}
	if ok := (&systemservice.SysDictDataService{}).IsHasSameValue(dictData.Value, dictData.DictId, &dictDataId); ok {
		response.NewError(nil).SetMsg("字典值已存在").Json(c)
		return
	}
	response.New(nil, (&systemservice.SysDictDataService{}).UpdateDictData(dictData)).Json(c)
}

// 字典数据删除
func (*DictController) DeleteData(c *gin.Context) {
	var ids []int
	if err := bind.BindIds(c, &ids); err != nil {
		response.NewError(err).Json(c)
		return
	}
	response.New(nil, (&systemservice.SysDictDataService{}).DeleteDictData(ids)).Json(c)
}
