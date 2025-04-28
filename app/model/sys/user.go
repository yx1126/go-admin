package model

import (
	"github.com/yx1126/go-admin/app/model"
	"github.com/yx1126/go-admin/app/util/datetime"
)

type SysUser struct {
	Id        int `gorm:"primaryKey;autoIncrement"`
	DeptId    int
	UserName  string
	NickName  string
	UserType  string `gorm:"default:00"`
	Email     string
	Phone     string
	Sex       string `gorm:"default:2"`
	Avatar    string
	Password  string
	LoginIp   string
	LoginDate datetime.Datetime
	Status    string `gorm:"default:1"`
	Remark    string
	model.BaseModel
}

func (SysUser) TableName() string {
	return "sys_user"
}
