package sysmodel

type SysRoleMenu struct {
	RoleId int `gorm:"primaryKey;AutoIncrement:false"`
	MenuId int `gorm:"primaryKey;AutoIncrement:false"`
}

func (*SysRoleMenu) TableName() string {
	return "sys_role_menu"
}
