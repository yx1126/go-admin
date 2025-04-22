package service

import (
	"github.com/yx1126/go-admin/DB"
	model "github.com/yx1126/go-admin/app/model/sys"
	"github.com/yx1126/go-admin/app/vo"
)

type SysDictDataService struct{}

func (*SysDictDataService) QueryDictDataAllList(dictId *int) ([]vo.DictDataListVo, error) {
	var dictDataList []vo.DictDataListVo
	query := DB.Gorm.Model(&model.SysDictData{}).Order("sort,id")
	if dictId != nil {
		query.Where("dict_id = ?", dictId)
	}
	result := query.Find(&dictDataList)
	return dictDataList, result.Error
}

func (*SysDictDataService) CreateDictData(dictData vo.CreateDictData) error {
	return DB.Gorm.Model(&model.SysDictData{}).Create(&model.SysDictData{
		DictId:    dictData.DictId,
		Sort:      dictData.Sort,
		Label:     dictData.Label,
		Value:     dictData.Value,
		Type:      dictData.Type,
		NodeType:  dictData.NodeType,
		CssClass:  dictData.CssClass,
		ListClass: dictData.ListClass,
		IsDefault: dictData.IsDefault,
		Status:    dictData.Status,
		Remark:    dictData.Remark,
	}).Error
}

func (*SysDictDataService) UpdateDictData(dictData vo.UpdateDictData) error {
	return DB.Gorm.Model(&model.SysDictData{}).Where("id = ?", dictData.Id).Updates(&model.SysDictData{
		DictId:    dictData.DictId,
		Sort:      dictData.Sort,
		Label:     dictData.Label,
		Value:     dictData.Value,
		Type:      dictData.Type,
		NodeType:  dictData.NodeType,
		CssClass:  dictData.CssClass,
		ListClass: dictData.ListClass,
		IsDefault: dictData.IsDefault,
		Status:    dictData.Status,
		Remark:    dictData.Remark,
	}).Error
}

func (*SysDictDataService) DeleteDictData(ids []int) error {
	return DB.Gorm.Model(&model.SysDictData{}).Delete(&model.SysDictData{}, ids).Error
}

func (*SysDictDataService) DictDataHasSameNameValue(label, value string, id *int) bool {
	var count int64
	query := DB.Gorm.Model(&model.SysDictData{}).Where("label = ?", label).Or("value = ?", value)
	if id != nil {
		query.Where("id != ?", id)
	}
	if result := query.Count(&count); result.Error != nil {
		return true
	}
	return count > 0
}
