package sysmodel

import (
	"github.com/yx1126/go-admin/app/model"
)

type SysMenu struct {
	ParentId int    `json:"parentId" gorm:"default:0"`
	Name     string `json:"name"`
	Title    string `json:"title"`
	// 0-目录；1-菜单；2-超链接；3-按钮
	Type uint   `json:"type" gorm:"default:0"`
	Link string `json:"link"`
	// 0-不缓存；1-缓存
	IsCache string `json:"isCache" gorm:"default:0"`
	Icon    string `json:"icon"`
	Path    string `json:"path"`
	// 0-否；1-是
	IsIframe   string `json:"isIframe" gorm:"default:0"`
	Component  string `json:"component"`
	Permission string `json:"permission"`
	Sort       int    `json:"sort" gorm:"default:0"`
	// 0-隐藏；1-显示
	Visible string `json:"visible" gorm:"default:1"`
	// 0-停用；1-正常
	Status string `json:"status" gorm:"default:1"`
	model.BaseModel
}

func (SysMenu) TableName() string {
	return "sys_menu"
}
