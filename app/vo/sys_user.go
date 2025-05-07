package vo

import (
	"github.com/yx1126/go-admin/app/model"
	sysmodel "github.com/yx1126/go-admin/app/model/sys"
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
	UserName  string `json:"userName"`
	LoginIp   string `json:"loginIp"`
	LoginDate string `json:"loginDate"`
	DeptId    *int   `json:"deptId"`
	DeptName  string `json:"deptName"`
	NickName  string `json:"nickName"`
	UserType  string `json:"userType"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Sex       string `json:"sex"`
	Avatar    string `json:"avatar"`
	Status    string `json:"status"`
	Remark    string `json:"remark"`
	model.BaseModel
}

// 详情
type UserInfoVo struct {
	PostIds *[]int `json:"postIds"`
	UserVo
}

func (UserVo) TableName() string {
	return (sysmodel.SysUser{}).TableName()
}

type baseCUVo struct {
	DeptId   *int   `json:"deptId"`
	PostIds  *[]int `json:"postIds"`
	DeptName string `json:"deptName"`
	NickName string `json:"nickName"`
	UserType string `json:"userType"`
	Email    string `json:"email" binding:"omitempty,email"`
	Phone    string `json:"phone"`
	Sex      string `json:"sex"`
	Status   string `json:"status"`
	Remark   string `json:"remark"`
}

type CreateUserVo struct {
	UserName string `json:"userName"`
	baseCUVo
}

type UpdateUserVo struct {
	Id int `json:"id"`
	baseCUVo
}

// 更新密码
type UpdatePwdVo struct {
	Id          int    `json:"id"`
	OldPassword string `json:"oldPassword" binding:"required"`
	Password    string `json:"password" binding:"required"`
	ConfirmPwd  string `json:"confirmPwd" binding:"required,eqfield=Password"`
}
