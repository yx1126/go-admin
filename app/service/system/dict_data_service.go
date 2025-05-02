package systemservice

import (
	"github.com/yx1126/go-admin/DB"
	sysmodel "github.com/yx1126/go-admin/app/model/sys"
	"github.com/yx1126/go-admin/app/service"
	"github.com/yx1126/go-admin/app/vo"
)

type SysDictDataService struct{}

func (*SysDictDataService) QueryDictDataList(dictId *int) ([]vo.DictDataListVo, error) {
	var dictDataList []vo.DictDataListVo
	query := DB.Gorm.Model(&sysmodel.SysDictData{}).
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
	result := DB.Gorm.Model(&sysmodel.SysDictData{}).
		Select("sys_dict_data.*", "t.type as dict_type", "t.node_type").
		Order("sort,id").
		Joins("LEFT JOIN sys_dict_type as t ON sys_dict_data.dict_id = t.id").
		Where("sys_dict_data.status = 1").
		Where("t.type = ?", dictType).
		Find(&dictDataList)
	return dictDataList, result.Error
}

func (*SysDictDataService) CreateDictData(dictData vo.CreateDictData) error {
	return DB.Gorm.Model(&sysmodel.SysDictData{}).Create(&sysmodel.SysDictData{
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
	return DB.Gorm.Model(&sysmodel.SysDictData{}).
		Scopes(service.UpdateOmitScope()).
		Where("id = ?", dictData.Id).
		Updates(&sysmodel.SysDictData{
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
	return DB.Gorm.Model(&sysmodel.SysDictData{}).Delete(&sysmodel.SysDictData{}, ids).Error
}

func (*SysDictDataService) DictDataHasSameNameValue(label, value string, dictId int, id *int) bool {
	var count int64
	query := DB.Gorm.Model(&sysmodel.SysDictData{}).
		Where(DB.Gorm.Where("label = ?", label).Or("value = ?", value)).
		Where("dict_id = ?", dictId)
	if id != nil {
		query.Where("id != ?", id)
	}
	if result := query.Count(&count); result.Error != nil {
		return true
	}
	return count > 0
}
