package service

import (
	"github.com/yx1126/go-admin/DB"
	model "github.com/yx1126/go-admin/app/model/sys"
	"github.com/yx1126/go-admin/app/util"
	"github.com/yx1126/go-admin/app/vo"
)

type MenuService struct{}

func (*MenuService) QueryMenuTree(menu vo.MenuQueryVo) ([]vo.MenuTreeVo, error) {
	menuList := make([]vo.MenuTreeVo, 0)
	query := DB.Gorm.Model(&model.SysMenu{}).Order("sys_menu.parent_id,sys_menu.sort,sys_menu.id")
	parentId := 0
	if menu.Title != "" {
		parentId = -1
		query.Where("title LIKE ?", "%"+menu.Title+"%")
	}
	if menu.Status != "" {
		parentId = -1
		query.Where("status = ?", menu.Status)
	}
	result := query.Find(&menuList)
	return util.ListToTree(menuList, parentId), result.Error
}

func (*MenuService) QueryMenuSelectTree() ([]vo.MenuTreeVo, error) {
	menuList := make([]vo.MenuTreeVo, 0)
	query := DB.Gorm.Model(&model.SysMenu{}).Order("sys_menu.parent_id,sys_menu.sort,sys_menu.id").Where("type != 3")
	result := query.Find(&menuList)
	return util.ListToTree(menuList, 0), result.Error
}

func (*MenuService) CreateMenu(menu vo.CreateMenuVo) error {
	return DB.Gorm.Model(&model.SysMenu{}).Create(&model.SysMenu{
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
	return DB.Gorm.Model(&model.SysMenu{}).Where("id = ?", menu.Id).Updates(&model.SysMenu{
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
	return DB.Gorm.Model(&model.SysMenu{}).Delete(&model.SysMenu{}, ids).Error
}

func (*MenuService) MenuHasChildren(parentId int) bool {
	var count int64
	result := DB.Gorm.Model(&model.SysMenu{}).Where("parent_id = ?", parentId).Count(&count)
	if result.Error != nil {
		count = 0
	}
	return count > 0
}

func (*MenuService) MenuHasSameName(name string, id *int) bool {
	var count int64
	query := DB.Gorm.Model(&model.SysMenu{}).Where("name = ?", name)
	if id != nil {
		query.Where("id != ?", id)
	}
	result := query.Count(&count)
	if result.Error != nil {
		count = 0
	}
	return count > 0
}
