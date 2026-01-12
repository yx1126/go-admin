package sysmodel

import "go-admin/app/model"

type SysRole struct {
	Name string `json:"name"`
	Key  string `json:"key"`
	Sort int    `json:"sort"`
	// 状态：0-停用；1-正常
	Status string `json:"status" gorm:"default:1"`
	Remark string `json:"remark"`
	model.BaseModel
}

func (*SysRole) TableName() string {
	return "sys_role"
}
