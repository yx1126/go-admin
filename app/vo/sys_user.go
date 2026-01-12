package vo

import (
	"go-admin/app/model"
	sysmodel "go-admin/app/model/sys"
	"go-admin/common/types"
)

// 列表查询
type UserParam struct {
	DeptId   string `json:"deptId" form:"deptId"`
	DeptName string `json:"deptName" form:"deptName"`
	UserName string `json:"userName" form:"userName"`
	NickName string `json:"nickName" form:"nickName"`
	Status   string `json:"status" form:"status"`
}

// 分页查询
type UserPagingParam struct {
	UserParam
	PagingVo
}

// 列表
type UserVo struct {
	UserName  string         `json:"userName"`
	LoginIp   string         `json:"loginIp"`
	LoginDate types.Datetime `json:"loginDate"`
	DeptId    *types.Long    `json:"deptId"`
	DeptName  string         `json:"deptName"`
	NickName  string         `json:"nickName"`
	UserType  string         `json:"userType"`
	Email     string         `json:"email"`
	Phone     string         `json:"phone"`
	Sex       string         `json:"sex"`
	Avatar    string         `json:"avatar"`
	Status    string         `json:"status"`
	Remark    string         `json:"remark"`
	model.BaseModel
}

type UserPwdVo struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
	Status   string `json:"status"`
	model.BaseModel
}

// 详情
type UserInfoVo struct {
	PostIds *[]types.Long `json:"postIds"`
	RoleIds *[]types.Long `json:"roleIds"`
	UserVo
}

func (UserVo) TableName() string {
	return (sysmodel.SysUser{}).TableName()
}

type baseCUVo struct {
	DeptId   *int          `json:"deptId,string" binding:"required"`
	PostIds  *[]types.Long `json:"postIds"`
	RoleIds  *[]types.Long `json:"roleIds"`
	NickName string        `json:"nickName" binding:"required"`
	UserType string        `json:"userType"`
	Email    string        `json:"email" binding:"omitempty,email"`
	Phone    string        `json:"phone" binding:"required"`
	Sex      string        `json:"sex"`
	Status   string        `json:"status"`
	Remark   string        `json:"remark"`
}

type CreateUserVo struct {
	UserName string `json:"userName" binding:"required"`
	baseCUVo
	CreatedBy string
}

type UpdateUserVo struct {
	BaseVo
	baseCUVo
	Avatar    string `json:"avatar"`
	UpdatedBy string
}

type UpdateUserLoginVo struct {
	BaseVo
	LoginIp   string         `json:"loginIp"`
	LoginDate types.Datetime `json:"loginDate"`
}

// 更新密码
type UpdatePwdVo struct {
	BaseVo
	OldPassword string `json:"oldPassword" binding:"required"`
	Password    string `json:"password" binding:"required"`
	ConfirmPwd  string `json:"confirmPwd" binding:"required,eqfield=Password"`
}
