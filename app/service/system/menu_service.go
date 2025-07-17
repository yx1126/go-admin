package systemservice

import (
	"github.com/yx1126/go-admin/DB"
	sysmodel "github.com/yx1126/go-admin/app/model/sys"
	"github.com/yx1126/go-admin/app/vo"
	"github.com/yx1126/go-admin/common/constant"
	"github.com/yx1126/go-admin/common/util"
)

type MenuService struct{}

// 查询菜单列表
func (*MenuService) QueryMenuList(menu vo.MenuParam) ([]vo.MenuTreeVo, error) {
	menuList := make([]vo.MenuTreeVo, 0)
	query := DB.Gorm.Model(&sysmodel.SysMenu{}).Order("sys_menu.parent_id,sys_menu.sort,sys_menu.id")
	if menu.Title != "" {
		query.Where("title LIKE ?", "%"+menu.Title+"%")
	}
	if menu.Status != "" {
		query.Where("status = ?", menu.Status)
	}
	result := query.Find(&menuList)
	return menuList, result.Error
}

// 查询下拉菜单树
func (*MenuService) QueryMenuSelectTree(status string) ([]vo.MenuTreeVo, error) {
	menuList := make([]vo.MenuTreeVo, 0)
	query := DB.Gorm.Model(&sysmodel.SysMenu{}).Order("sys_menu.parent_id,sys_menu.sort,sys_menu.id")
	if status != "" {
		query.Where("status = ?", status)
	}
	result := query.Find(&menuList)
	return util.ListToTree(menuList, 0), result.Error
}

// 创建菜单
func (*MenuService) CreateMenu(menu vo.CreateMenuVo) error {
	return DB.Gorm.Model(&sysmodel.SysMenu{}).Create(&sysmodel.SysMenu{
		ParentId:   menu.ParentId,
		Name:       menu.Name,
		Type:       menu.Type,
		Link:       menu.Link,
		Title:      menu.Title,
		IsCache:    menu.IsCache,
		Icon:       menu.Icon,
		Path:       menu.Path,
		IsIframe:   menu.IsIframe,
		Component:  menu.Component,
		Permission: menu.Permission,
		Sort:       menu.Sort,
		Visible:    menu.Visible,
	}).Error
}

// 更新菜单
func (*MenuService) UpdateMenu(menu vo.UpdateMenuVo) error {
	return DB.Gorm.Model(&sysmodel.SysMenu{}).Where("id = ?", menu.Id).Updates(&sysmodel.SysMenu{
		ParentId:   menu.ParentId,
		Name:       menu.Name,
		Type:       menu.Type,
		Link:       menu.Link,
		Title:      menu.Title,
		IsCache:    menu.IsCache,
		Icon:       menu.Icon,
		Path:       menu.Path,
		IsIframe:   menu.IsIframe,
		Component:  menu.Component,
		Permission: menu.Permission,
		Sort:       menu.Sort,
		Visible:    menu.Visible,
		Status:     menu.Status,
	}).Error
}

// 删除菜单
func (*MenuService) DeleteMenus(ids []int) error {
	return DB.Gorm.Model(&sysmodel.SysMenu{}).Delete(&sysmodel.SysMenu{}, ids).Error
}

// 校验存在子元素
func (*MenuService) IsHasChildren(parentId int) bool {
	var count int64
	DB.Gorm.Model(&sysmodel.SysMenu{}).Where("parent_id = ?", parentId).Count(&count)
	return count > 0
}

// 校验菜单页面名称
func (*MenuService) IsHasSameName(name string, id *int) bool {
	var count int64
	query := DB.Gorm.Model(&sysmodel.SysMenu{}).Where("name = ?", name)
	if id != nil {
		query.Where("id != ?", id)
	}
	query.Count(&count)
	return count > 0
}

func (*MenuService) QueryAuthMenuList(id int) ([]vo.MenuTreeVo, error) {
	menuList := make([]vo.MenuTreeVo, 0)
	query := DB.Gorm.Model(&sysmodel.SysMenu{}).Distinct("sys_menu.*").
		Order("sys_menu.parent_id,sys_menu.sort,sys_menu.id")
	if id != constant.ADMIN_ID {
		query.Joins("LEFT JOIN sys_role_menu ON sys_menu.id = sys_role_menu.menu_id").
			Joins("LEFT JOIN sys_role ON sys_role_menu.role_id = sys_role.id").
			Joins("LEFT JOIN sys_user_role ON sys_role.id = sys_user_role.role_id").
			Where("sys_user_role.user_id = ? AND sys_role.status = ?", id, constant.STATUS)
	}
	query.Where("sys_menu.status = ? AND sys_menu.type IN ?", constant.STATUS, []string{"0", "1", "2"})
	result := query.Find(&menuList)
	return menuList, result.Error
}

// 根据用户id查询菜单权限perms
func (s *MenuService) QueryPermsByUserId(userId int) []string {
	perms := make([]string, 0)
	// 超级管理员拥有所有权限
	if userId == constant.ADMIN_ID {
		perms = append(perms, "*:*:*")
	} else {
		DB.Gorm.Model(&sysmodel.SysMenu{}).
			Distinct("sys_menu.permission").
			Joins("JOIN sys_role_menu ON sys_menu.id = sys_role_menu.menu_id").
			Joins("JOIN sys_role ON sys_role_menu.role_id = sys_role.id").
			Joins("JOIN sys_user_role ON sys_role.id = sys_user_role.role_id").
			Where("sys_user_role.user_id = ? AND sys_menu.status = ? AND sys_menu.type = ?", userId, constant.STATUS, 3).
			Pluck("sys_menu.permission", &perms)
	}
	return perms
}
