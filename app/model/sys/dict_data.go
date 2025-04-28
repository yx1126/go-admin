package model

import (
	"github.com/yx1126/go-admin/app/model"
)

type SysDictData struct {
	Id        int `gorm:"primaryKey;autoIncrement"`
	DictId    int
	Sort      int `gorm:"default:0"`
	Label     string
	Value     string
	Type      string
	CssClass  string
	ListClass string
	// 0-否；1-是
	IsDefault string `gorm:"default:0"`
	// 0-停用；1-正常
	Status string `gorm:"default:1"`
	Remark string
	model.BaseModel
}

func (SysDictData) TableName() string {
	return "sys_dict_data"
}
