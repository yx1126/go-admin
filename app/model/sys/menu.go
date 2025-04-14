package model

import (
	"github.com/yx1126/go-admin/app/model"
)

type SysMenu struct {
	Id         uint `gorm:"primaryKey;autoIncrement"`
	ParentId   uint `gorm:"default:0"`
	Name       string
	Title      string
	Type       uint `gorm:"default:0"`
	Link       string
	IsCache    string `gorm:"default:0"`
	Icon       string
	Path       string
	IsIframe   string `gorm:"default:0"`
	Component  string
	Permission string
	Sort       int    `gorm:"default:0"`
	Visible    string `gorm:"default:1"`
	Status     string `gorm:"default:1"`
	model.BaseModel
}

func (SysMenu) TableName() string {
	return "sys_menu"
}
