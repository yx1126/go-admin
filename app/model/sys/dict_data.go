package sysmodel

import (
	"go-admin/app/model"
)

type SysDictData struct {
	DictId    int    `json:"dictId"`
	Sort      int    `json:"sort" gorm:"default:0"`
	Label     string `json:"label"`
	Value     string `json:"value"`
	Type      string `json:"type"`
	CssClass  string `json:"cssClass"`
	ListClass string `json:"listClass"`
	// 0-否；1-是
	IsDefault string `json:"isDefault" gorm:"default:0"`
	// 0-停用；1-正常
	Status string `json:"status" gorm:"default:1"`
	Remark string `json:"remark"`
	model.BaseModel
}

func (SysDictData) TableName() string {
	return "sys_dict_data"
}
