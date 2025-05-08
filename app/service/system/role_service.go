package systemservice

import (
	"github.com/yx1126/go-admin/DB"
	sysmodel "github.com/yx1126/go-admin/app/model/sys"
	"github.com/yx1126/go-admin/app/service"
	"github.com/yx1126/go-admin/app/vo"
)

type RoleService struct{}

func (*RoleService) QueryRoleList(param vo.RoleParam) (vo.PagingBackVo[vo.RoleVo], error) {
	var roleList []vo.RoleVo
	var count int64
	query := DB.Gorm.Model(&sysmodel.SysRole{}).Scopes(service.SortScope)
	if param.Name != "" {
		query.Where("name LIKE ?", "%"+param.Name+"%")
	}
	if param.Key != "" {
		query.Where("key LIKE ?", "%"+param.Key+"%")
	}
	if param.Status != "" {
		query.Where("status = ?", param.Status)
	}
	result := query.Count(&count).
		Scopes(service.PagingScope(param.Page, param.Size)).
		Find(&roleList)
	return vo.PagingBackVo[vo.RoleVo]{Data: roleList, Count: int(count)}, result.Error
}

func (*RoleService) CreateRole(role vo.CreateRoleVo) error {
	return DB.Gorm.Model(&sysmodel.SysRole{}).Create(&sysmodel.SysRole{
		Name:   role.Name,
		Key:    role.Key,
		Sort:   role.Sort,
		Status: role.Status,
		Remark: role.Remark,
	}).Error
}

func (*RoleService) UpdateRole(role vo.UpdateRoleVo) error {
	return DB.Gorm.Model(&sysmodel.SysRole{}).Where("id = ?", role.Id).Updates(&sysmodel.SysRole{
		Name:   role.Name,
		Key:    role.Key,
		Sort:   role.Sort,
		Status: role.Status,
		Remark: role.Remark,
	}).Error
}

func (*RoleService) DeleteRole(ids []int) error {
	return DB.Gorm.Model(&sysmodel.SysRole{}).Delete(&sysmodel.SysRole{}, ids).Error
}

func (*RoleService) IsHasSameName(name string, id *int) bool {
	var count int64
	query := DB.Gorm.Model(&sysmodel.SysRole{}).Where("name = ?", name)
	if id != nil {
		query.Where("id = ?", id)
	}
	query.Count(&count)
	return count > 0
}

func (*RoleService) IsHasSameKey(key string, id *int) bool {
	var count int64
	query := DB.Gorm.Model(&sysmodel.SysRole{}).Where("key = ?", key)
	if id != nil {
		query.Where("id = ?", id)
	}
	query.Count(&count)
	return count > 0
}
