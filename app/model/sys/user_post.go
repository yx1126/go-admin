package sysmodel

type SysUserPost struct {
	UserId int `gorm:"primaryKey"`
	PostId int `gorm:"primaryKey"`
}

func (*SysUserPost) TableName() string {
	return "sys_user_post"
}
