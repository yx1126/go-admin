package service

import (
	"github.com/yx1126/go-admin/DB"
	model "github.com/yx1126/go-admin/app/model/sys"
	"github.com/yx1126/go-admin/app/vo"
)

type SysDictTypeService struct{}

func (*SysDictTypeService) QueryDictTypeAllList() ([]vo.DictTypeListVo, error) {
	var dictTypeList []vo.DictTypeListVo
	result := DB.Gorm.Model(&model.SysDictType{}).Order("updated_at,created_at,id").Find(&dictTypeList)
	return dictTypeList, result.Error
}

func (*SysDictTypeService) CreateDictType(dictType vo.CreateDictType) error {
	return DB.Gorm.Model(&model.SysDictType{}).Create(&model.SysDictType{
		Name:     dictType.Name,
		Type:     dictType.Type,
		NodeType: dictType.NodeType,
		Status:   dictType.Status,
		Remark:   dictType.Remark,
	}).Error
}

func (*SysDictTypeService) UpdateDictType(dictType vo.UpdateDictType) error {
	return DB.Gorm.Model(&model.SysDictType{}).Where("id = ?", dictType.Id).Updates(&model.SysDictType{
		Name:     dictType.Name,
		Type:     dictType.Type,
		NodeType: dictType.NodeType,
		Status:   dictType.Status,
		Remark:   dictType.Remark,
	}).Error
}

func (*SysDictTypeService) DeleteDictType(ids []int) error {
	return DB.Gorm.Model(&model.SysDictType{}).Delete(&model.SysDictType{}, ids).Error
}

func (*SysDictTypeService) DictTypeHasSameType(typeName string, id *int) bool {
	var count int64
	query := DB.Gorm.Model(&model.SysDictType{}).Where("type = ?", typeName)
	if id != nil {
		query.Where("id != ?", id)
	}
	if result := query.Count(&count); result.Error != nil {
		return true
	}
	return count > 0
}
