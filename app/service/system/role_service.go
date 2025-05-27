package systemservice

import (
	"github.com/yx1126/go-admin/DB"
	sysmodel "github.com/yx1126/go-admin/app/model/sys"
	"github.com/yx1126/go-admin/app/service"
	"github.com/yx1126/go-admin/app/vo"
	"github.com/yx1126/go-admin/common/constant"
	"github.com/yx1126/go-admin/common/types"
	"github.com/yx1126/go-admin/common/util"
)

type RoleService struct{}

// 分页查询角色
func (*RoleService) QueryRoleList(param vo.RoleParam) (vo.PagingBackVo[vo.RoleVo], error) {
	var roleList []vo.RoleVo
	var count int64
	query := DB.Gorm.Model(&sysmodel.SysRole{}).Order("sort,updated_at DESC,created_at DESC")
	if param.Name != "" {
		query.Where("name LIKE ?", "%"+param.Name+"%")
	}
	if param.Key != "" {
		query.Where("sys_role.key LIKE ?", "%"+param.Key+"%")
	}
	if param.Status != "" {
		query.Where("status = ?", param.Status)
	}
	result := query.Count(&count).
		Scopes(service.PagingScope(param.Page, param.Size)).
		Find(&roleList)
	return vo.PagingBackVo[vo.RoleVo]{Data: roleList, Count: int(count)}, result.Error
}

// 查询全部角色
func (*RoleService) QueryRoleAllList() ([]vo.RoleVo, error) {
	var roleList []vo.RoleVo
	result := DB.Gorm.Model(&sysmodel.SysRole{}).Where("status = ?", constant.STATUS).Find(&roleList)
	return roleList, result.Error
}

// 查询角色信息
func (*RoleService) QueryRoleInfo(id int) (*vo.RoleInfoVo, error) {
	var role vo.RoleInfoVo
	if err := DB.Gorm.Model(&sysmodel.SysRole{}).Where("id = ?", id).First(&role.RoleVo).Error; err != nil {
		return nil, err
	}
	if err := DB.Gorm.Model(&sysmodel.SysRoleMenu{}).
		Select("menu_id").
		Where("role_id = ?", id).
		Pluck("menu_id", &role.MenuIds).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

// 创建角色
func (*RoleService) CreateRole(role vo.CreateRoleVo) error {
	tx := DB.Gorm.Begin()
	roleInfo := sysmodel.SysRole{
		Name:   role.Name,
		Key:    role.Key,
		Sort:   role.Sort,
		Status: role.Status,
		Remark: role.Remark,
	}
	if err := tx.Model(&sysmodel.SysRole{}).Create(&roleInfo).Error; err != nil {
		tx.Rollback()
		return err
	}
	// 插入菜单
	if role.MenuIds != nil && len(*role.MenuIds) > 0 {
		roleMenus := util.Map(*role.MenuIds, func(item types.Long, _ int) sysmodel.SysRoleMenu {
			return sysmodel.SysRoleMenu{
				RoleId: roleInfo.Id,
				MenuId: item.Val,
			}
		})
		if err := tx.Model(&sysmodel.SysRoleMenu{}).Create(&roleMenus).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit().Error
}

// 更新角色
func (*RoleService) UpdateRole(role vo.UpdateRoleVo) error {
	tx := DB.Gorm.Begin()
	if err := tx.Model(&sysmodel.SysRole{}).Where("id = ?", role.Id).Updates(&sysmodel.SysRole{
		Name:   role.Name,
		Key:    role.Key,
		Sort:   role.Sort,
		Status: role.Status,
		Remark: role.Remark,
	}).Error; err != nil {
		tx.Rollback()
		return err
	}
	// 删除菜单
	if err := tx.Model(&sysmodel.SysRoleMenu{}).Where("role_id = ?", role.Id).Delete(&sysmodel.SysRoleMenu{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	// 插入菜单
	if role.MenuIds != nil && len(*role.MenuIds) > 0 {
		roleMenus := util.Map(*role.MenuIds, func(item types.Long, _ int) sysmodel.SysRoleMenu {
			return sysmodel.SysRoleMenu{
				RoleId: role.Id,
				MenuId: item.Val,
			}
		})
		if err := tx.Model(&sysmodel.SysRoleMenu{}).Create(&roleMenus).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit().Error
}

// 删除角色
func (*RoleService) DeleteRole(ids []int) error {
	tx := DB.Gorm.Begin()
	if err := DB.Gorm.Model(&sysmodel.SysRole{}).Delete(&sysmodel.SysRole{}, ids).Error; err != nil {
		tx.Rollback()
		return err
	}
	// 删除菜单
	if err := tx.Model(&sysmodel.SysRoleMenu{}).Where("role_id in ?", ids).Delete(&sysmodel.SysRoleMenu{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	// 删除菜单
	return tx.Commit().Error
}

// 校验角色名称
func (*RoleService) IsHasSameName(name string, id *int) bool {
	var count int64
	query := DB.Gorm.Model(&sysmodel.SysRole{}).Where("name = ?", name)
	if id != nil {
		query.Where("id != ?", id)
	}
	query.Count(&count)
	return count > 0
}

// 校验角色权限字符
func (*RoleService) IsHasSameKey(key string, id *int) bool {
	var count int64
	query := DB.Gorm.Model(&sysmodel.SysRole{}).Where("key = ?", key)
	if id != nil {
		query.Where("id != ?", id)
	}
	query.Count(&count)
	return count > 0
}
