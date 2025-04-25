package model

import (
	"github.com/yx1126/go-admin/app/model"
)

type SysDictType struct {
	Id   uint `gorm:"primaryKey;autoIncrement"`
	Name string
	Type string
	// 0-文本；1-标签
	NodeType string `gorm:"default:0"`
	// 0-停用；1-正常
	Status string `gorm:"default:1"`
	Remark string
	model.BaseModel
}

func (SysDictType) TableName() string {
	return "sys_dict_type"
}
