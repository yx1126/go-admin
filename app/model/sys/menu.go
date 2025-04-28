package model

import (
	"github.com/yx1126/go-admin/app/model"
)

type SysMenu struct {
	Id       int `gorm:"primaryKey;autoIncrement"`
	ParentId int `gorm:"default:0"`
	Name     string
	Title    string
	// 0-目录；1-菜单；2-超链接；3-按钮
	Type uint `gorm:"default:0"`
	Link string
	// 0-不缓存；1-缓存
	IsCache string `gorm:"default:0"`
	Icon    string
	Path    string
	// 0-否；1-是
	IsIframe   string `gorm:"default:0"`
	Component  string
	Permission string
	Sort       int `gorm:"default:0"`
	// 0-隐藏；1-显示
	Visible string `gorm:"default:1"`
	// 0-停用；1-正常
	Status string `gorm:"default:1"`
	model.BaseModel
}

func (SysMenu) TableName() string {
	return "sys_menu"
}
