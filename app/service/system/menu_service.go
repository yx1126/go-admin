package systemservice

import (
	"github.com/yx1126/go-admin/DB"
	sysmodel "github.com/yx1126/go-admin/app/model/sys"
	"github.com/yx1126/go-admin/app/util"
	"github.com/yx1126/go-admin/app/vo"
)

type MenuService struct{}

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

func (*MenuService) QueryMenuSelectTree() ([]vo.MenuTreeVo, error) {
	menuList := make([]vo.MenuTreeVo, 0)
	query := DB.Gorm.Model(&sysmodel.SysMenu{}).Order("sys_menu.parent_id,sys_menu.sort,sys_menu.id").Where("type != 3")
	result := query.Find(&menuList)
	return util.ListToTree(menuList, 0), result.Error
}

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

func (*MenuService) DeleteMenus(ids []int) error {
	return DB.Gorm.Model(&sysmodel.SysMenu{}).Delete(&sysmodel.SysMenu{}, ids).Error
}

func (*MenuService) MenuHasChildren(parentId int) bool {
	var count int64
	result := DB.Gorm.Model(&sysmodel.SysMenu{}).Where("parent_id = ?", parentId).Count(&count)
	if result.Error != nil {
		count = 0
	}
	return count > 0
}

func (*MenuService) MenuHasSameName(name string, id *int) bool {
	var count int64
	query := DB.Gorm.Model(&sysmodel.SysMenu{}).Where("name = ?", name)
	if id != nil {
		query.Where("id != ?", id)
	}
	if result := query.Count(&count); result.Error != nil {
		count = 0
	}
	return count > 0
}
