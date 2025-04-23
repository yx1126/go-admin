package model

import (
	"github.com/yx1126/go-admin/app/model"
)

type SysDictData struct {
	Id        uint `gorm:"primaryKey;autoIncrement"`
	DictId    uint
	Sort      int `gorm:"default:0"`
	Label     string
	Value     string
	Type      string
	CssClass  string
	ListClass string
	IsDefault string `gorm:"default:0"`
	Status    string `gorm:"default:1"`
	Remark    string
	model.BaseModel
}

func (SysDictData) TableName() string {
	return "sys_dict_data"
}
