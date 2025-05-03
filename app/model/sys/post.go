package sysmodel

import "github.com/yx1126/go-admin/app/model"

type SysPost struct {
	Code   string `json:"code"`
	Name   string `json:"name"`
	Sort   int    `json:"sort" gorm:"default:0"`
	Status string `json:"status" gorm:"default:1"`
	Remark string `json:"remark"`
	model.BaseModel
}

func (SysPost) TableName() string {
	return "sys_post"
}
