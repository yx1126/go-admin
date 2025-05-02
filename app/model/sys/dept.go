package sysmodel

import "github.com/yx1126/go-admin/app/model"

type SysDept struct {
	ParentId int    `json:"parentId" gorm:"default:0"`
	Name     string `json:"name"`
	Sort     int    `json:"sort" gorm:"default:0"`
	LeaderId *int   `json:"leaderId"`
	// 0-停用；1-启用
	Status string `json:"status" gorm:"default:1"`
	model.BaseModel
}

func (SysDept) TableName() string {
	return "sys_dept"
}
