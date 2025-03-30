package service

import (
	"github.com/yx1126/go-admin/DB"
	model "github.com/yx1126/go-admin/app/model/sys"
	"github.com/yx1126/go-admin/app/vo"
)

type MenuService struct{}

func (m *MenuService) CreateMenu(menu vo.CreateMenuVo) error {
	result := DB.Gorm.Model(&model.SysMenu{}).Create(&model.SysMenu{
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
	})
	return result.Error
}

func (m *MenuService) QueryMenuTree(menu vo.MenuQueryVo) ([]vo.MenuTreeVo, error) {
	menuList := make([]vo.MenuTreeVo, 0)
	query := DB.Gorm.Model(&model.SysMenu{}).Order("sys_menu.parent_id,sys_menu.sort,sys_menu.id")
	if menu.Name != "" {
		query.Where("name LIKE ?", "&"+menu.Name+"%")
	}
	if menu.Status != "" {
		query.Where("status = ?", menu.Status)
	}
	result := query.Find(&menuList)
	return menuList, result.Error
}
