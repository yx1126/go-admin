package service

import (
	"github.com/yx1126/go-admin/db"
	model "github.com/yx1126/go-admin/app/model/sys"
	"github.com/yx1126/go-admin/app/service"
	"github.com/yx1126/go-admin/app/vo"
)

type SysDictDataService struct{}

func (*SysDictDataService) QueryDictDataList(dictId *int) ([]vo.DictDataListVo, error) {
	var dictDataList []vo.DictDataListVo
	query := db.Gorm.Model(&model.SysDictData{}).
		Select("sys_dict_data.*", "t.type as dict_type", "t.node_type").
		Order("sort,id").
		Joins("LEFT JOIN sys_dict_type as t ON sys_dict_data.dict_id = t.id")
	if dictId != nil {
		query.Where("dict_id = ?", dictId)
	}
	result := query.Find(&dictDataList)
	return dictDataList, result.Error
}

func (*SysDictDataService) QueryDictDataListByType(dictType string) ([]vo.DictDataListVo, error) {
	var dictDataList []vo.DictDataListVo
	result := db.Gorm.Model(&model.SysDictData{}).
		Select("sys_dict_data.*", "t.type as dict_type", "t.node_type").
		Order("sort,id").
		Joins("LEFT JOIN sys_dict_type as t ON sys_dict_data.dict_id = t.id").
		Where("t.type = ?", dictType).
		Find(&dictDataList)
	return dictDataList, result.Error
}

func (*SysDictDataService) CreateDictData(dictData vo.CreateDictData) error {
	return db.Gorm.Model(&model.SysDictData{}).Create(&model.SysDictData{
		DictId:    dictData.DictId,
		Sort:      dictData.Sort,
		Label:     dictData.Label,
		Value:     dictData.Value,
		Type:      dictData.Type,
		CssClass:  dictData.CssClass,
		ListClass: dictData.ListClass,
		IsDefault: dictData.IsDefault,
		Status:    dictData.Status,
		Remark:    dictData.Remark,
	}).Error
}

func (*SysDictDataService) UpdateDictData(dictData vo.UpdateDictData) error {
	return db.Gorm.Model(&model.SysDictData{}).
		Scopes(service.UpdateOmitScope()).
		Where("id = ?", dictData.Id).
		Updates(&model.SysDictData{
			DictId:    dictData.DictId,
			Sort:      dictData.Sort,
			Label:     dictData.Label,
			Value:     dictData.Value,
			Type:      dictData.Type,
			CssClass:  dictData.CssClass,
			ListClass: dictData.ListClass,
			IsDefault: dictData.IsDefault,
			Status:    dictData.Status,
			Remark:    dictData.Remark,
		}).Error
}

func (*SysDictDataService) DeleteDictData(ids []int) error {
	return db.Gorm.Model(&model.SysDictData{}).Delete(&model.SysDictData{}, ids).Error
}

func (*SysDictDataService) DictDataHasSameNameValue(label, value string, dictId uint, id *int) bool {
	var count int64
	query := db.Gorm.Model(&model.SysDictData{}).
		Where(db.Gorm.Where("label = ?", label).Or("value = ?", value)).
		Where("dict_id = ?", dictId)
	if id != nil {
		query.Where("id != ?", id)
	}
	if result := query.Count(&count); result.Error != nil {
		return true
	}
	return count > 0
}
