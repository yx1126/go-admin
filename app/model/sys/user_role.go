package sysmodel

type SysUserRole struct {
	UserId int `gorm:"primaryKey"`
	RoleId int `gorm:"primaryKey"`
}

func (*SysUserRole) TableName() string {
	return "sys_role_menu"
}
