package sysmodel

type SysUserPost struct {
	UserId int `gorm:"primaryKey;AutoIncrement:false"`
	PostId int `gorm:"primaryKey;AutoIncrement:false"`
}

func (*SysUserPost) TableName() string {
	return "sys_user_post"
}
