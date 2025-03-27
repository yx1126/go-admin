package model

import (
	"time"

	"github.com/yx1126/go-admin/app/model"
)

type SysUser struct {
	UserId      int `gorm:"primaryKey;autoIncrement"`
	DeptId      int
	UserName    string
	NickName    string
	UserType    string `gorm:"default:00"`
	Email       string
	Phonenumber string
	Sex         string `gorm:"default:2"`
	Avatar      string
	Password    string
	Status      string `gorm:"default:0"`
	LoginIp     string
	LoginDate   time.Time
	Remark      string
	model.BaseModel
}

func (SysUser) TableName() string {
	return "sys_user"
}
