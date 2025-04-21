package model

import (
	"github.com/yx1126/go-admin/app/model"
)

type SysDictType struct {
	Id     uint `gorm:"primaryKey;autoIncrement"`
	Name   string
	Type   string
	Status string `gorm:"default:1"`
	Remark string
	model.BaseModel
}

func (SysDictType) TableName() string {
	return "sys_dict_type"
}
