package model

import (
	"github.com/yx1126/go-admin/app/model"
)

type SysMenu struct {
	Id         int `gorm:"primaryKey;autoIncrement"`
	ParentId   *int
	Name       string
	Type       int
	Link       string
	Title      string
	IsCache    int `gorm:"default:0"`
	Icon       string
	Path       string
	IsIframe   int `gorm:"default:0"`
	Component  string
	Permission string
	Sort       *int
	Visible    int `gorm:"default:0"`
	model.BaseModel
}

func (SysMenu) TableName() string {
	return "sys_menu"
}
