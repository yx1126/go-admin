package model

import (
	"github.com/yx1126/go-admin/app/model"
)

type SysMenu struct {
	Id         uint `gorm:"primaryKey;autoIncrement"`
	ParentId   *uint
	Name       string
	Type       string
	Link       string
	Title      string
	IsCache    string `gorm:"default:0"`
	Icon       string
	Path       string
	IsIframe   string `gorm:"default:0"`
	Component  string
	Permission string
	Sort       *int
	Visible    string `gorm:"default:0"`
	Status     string `gorm:"default:0"`
	model.BaseModel
}

func (SysMenu) TableName() string {
	return "sys_menu"
}
