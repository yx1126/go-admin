package sysmodel

type SysRoleMenu struct {
	RoleId int `gorm:"primaryKey"`
	MenuId int `gorm:"primaryKey"`
}

func (*SysRoleMenu) TableName() string {
	return "sys_role_menu"
}
