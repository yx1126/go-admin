package service

import (
	"github.com/yx1126/go-admin/DB"
	model "github.com/yx1126/go-admin/app/model/sys"
	"github.com/yx1126/go-admin/app/vo"
)

type SysDictTypeService struct{}

func (*SysDictTypeService) QueryDictTypeAllList() ([]vo.DictTypeListVo, error) {
	var dictTypeList []vo.DictTypeListVo
	result := DB.Gorm.Model(&model.SysDictType{}).Order("sys_dict_type.id").Find(&dictTypeList)
	return dictTypeList, result.Error
}

func (*SysDictTypeService) CreateDictType(dictType vo.CreateDictType) error {
	return DB.Gorm.Model(&model.SysDictType{}).Create(&model.SysDictType{
		Name:   dictType.Name,
		Type:   dictType.Type,
		Status: dictType.Status,
		Remark: dictType.Remark,
	}).Error
}

func (*SysDictTypeService) UpdateDictType(dictType vo.UpdateDictType) error {
	return DB.Gorm.Model(&model.SysDictType{}).Updates(&model.SysDictType{
		Name:   dictType.Name,
		Type:   dictType.Type,
		Status: dictType.Status,
		Remark: dictType.Remark,
	}).Where("id = ?", dictType.Id).Error
}

func (*SysDictTypeService) DeleteDictType(ids []int) error {
	return DB.Gorm.Model(&model.SysDictType{}).Delete(&model.SysDictType{}, ids).Error
}
