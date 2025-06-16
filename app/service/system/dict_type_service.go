package systemservice

import (
	"github.com/yx1126/go-admin/DB"
	sysmodel "github.com/yx1126/go-admin/app/model/sys"
	"github.com/yx1126/go-admin/app/vo"
)

type SysDictTypeService struct{}

// 字典查询
func (*SysDictTypeService) QueryDictTypeAllList() ([]vo.DictTypeListVo, error) {
	var dictTypeList []vo.DictTypeListVo
	result := DB.Gorm.Model(&sysmodel.SysDictType{}).Order("updated_at,created_at,id").Find(&dictTypeList)
	return dictTypeList, result.Error
}

// 字典查询
func (*SysDictTypeService) QueryDictTypeById(id int) (vo.DictTypeListVo, error) {
	var dictTypeList vo.DictTypeListVo
	result := DB.Gorm.Model(&sysmodel.SysDictType{}).Where("id = ?", id).First(&dictTypeList)
	return dictTypeList, result.Error
}

// 创建字典
func (*SysDictTypeService) CreateDictType(dictType vo.CreateDictType) error {
	return DB.Gorm.Model(&sysmodel.SysDictType{}).Create(&sysmodel.SysDictType{
		Name:     dictType.Name,
		Type:     dictType.Type,
		NodeType: dictType.NodeType,
		Status:   dictType.Status,
		Remark:   dictType.Remark,
	}).Error
}

// 更新字典
func (*SysDictTypeService) UpdateDictType(dictType vo.UpdateDictType) error {
	return DB.Gorm.Model(&sysmodel.SysDictType{}).Where("id = ?", dictType.Id).Updates(&sysmodel.SysDictType{
		Name:     dictType.Name,
		Type:     dictType.Type,
		NodeType: dictType.NodeType,
		Status:   dictType.Status,
		Remark:   dictType.Remark,
	}).Error
}

// 删除字典
func (*SysDictTypeService) DeleteDictType(ids []int) error {
	return DB.Gorm.Model(&sysmodel.SysDictType{}).Delete(&sysmodel.SysDictType{}, ids).Error
}

// 校验字典类型
func (*SysDictTypeService) IsHasSameType(typeName string, id *int) bool {
	var count int64
	query := DB.Gorm.Model(&sysmodel.SysDictType{}).Where("type = ?", typeName)
	if id != nil {
		query.Where("id != ?", id)
	}
	query.Count(&count)
	return count > 0
}
