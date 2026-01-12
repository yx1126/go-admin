package systemcontroller

import (
	systemservice "go-admin/app/service/system"
	"go-admin/app/vo"
	bind "go-admin/common/should_bind"
	"go-admin/response"

	"github.com/gin-gonic/gin"
)

type DictController struct{}

// 字典类型查询
func (*DictController) QueryAllList(c *gin.Context) {
	response.New((&systemservice.SysDictTypeService{}).QueryDictTypeAllList()).Json(c)
}

// 字典类型新增
func (*DictController) Create(c *gin.Context) {
	var dictType vo.CreateDictType
	err := bind.ShouldBindJSON(c, &dictType)
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
	err := bind.ShouldBindJSON(c, &dictType)
	if err != nil {
		response.NewError(err).Json(c)
		return
	}
	if ok := (&systemservice.SysDictTypeService{}).IsHasSameType(dictType.Type, &dictType.Id); ok {
		response.NewError(nil).SetMsg("字典类型已存在").Json(c)
		return
	}
	response.New(nil, (&systemservice.SysDictTypeService{}).UpdateDictType(dictType)).Json(c)
}

// 字典类型删除
func (*DictController) Delete(c *gin.Context) {
	var ids []int
	err := bind.BindIds(c, &ids)
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
	response.New(nil, err).SetPaging(paging).Json(c)
}

// 根据字典类型查询字典数据
func (*DictController) QueryDictDataListByType(c *gin.Context) {
	response.New((&systemservice.SysDictDataService{}).QueryDictDataListByType(c.Param("dictType"))).Json(c)
}

// 字典数据新增
func (*DictController) CreateData(c *gin.Context) {
	var dictData vo.CreateDictData
	err := bind.ShouldBindJSON(c, &dictData)
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
	err := bind.ShouldBindJSON(c, &dictData)
	if err != nil {
		response.NewError(err).Json(c)
		return
	}
	if ok := (&systemservice.SysDictDataService{}).IsHasSameName(dictData.Label, dictData.DictId, &dictData.Id); ok {
		response.NewError(nil).SetMsg("字典名称已存在").Json(c)
		return
	}
	if ok := (&systemservice.SysDictDataService{}).IsHasSameValue(dictData.Value, dictData.DictId, &dictData.Id); ok {
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
