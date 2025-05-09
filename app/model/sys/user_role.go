package sysmodel

type SysUserRole struct {
	UserId int `gorm:"primaryKey;AutoIncrement:false"`
	RoleId int `gorm:"primaryKey;AutoIncrement:false"`
}

func (*SysUserRole) TableName() string {
	return "sys_user_role"
}
