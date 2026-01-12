package sysmodel

import (
	"go-admin/app/model"
)

type SysDictType struct {
	Name string `json:"name"`
	Type string `json:"type"`
	// 0-文本；1-标签
	NodeType string `json:"nodeType" gorm:"default:0"`
	// 0-停用；1-正常
	Status string `json:"status" gorm:"default:1"`
	Remark string `json:"remark"`
	model.BaseModel
}

func (SysDictType) TableName() string {
	return "sys_dict_type"
}
