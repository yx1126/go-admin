package sysmodel

import (
	"github.com/yx1126/go-admin/app/model"
	"github.com/yx1126/go-admin/common/types"
)

type SysUser struct {
	DeptId    *int           `json:"deptId"`
	UserName  string         `json:"userName"`
	NickName  string         `json:"nickName"`
	UserType  string         `json:"userType" gorm:"default:00"`
	Email     string         `json:"email"`
	Phone     string         `json:"phone"`
	Sex       string         `json:"sex" gorm:"default:2"`
	Avatar    string         `json:"avatar"`
	Password  string         `json:"-"`
	LoginIp   string         `json:"loginIp"`
	LoginDate types.Datetime `json:"loginDate"`
	Status    string         `json:"status" gorm:"default:1"`
	Remark    string         `json:"remark"`
	model.BaseModel
}

func (SysUser) TableName() string {
	return "sys_user"
}
