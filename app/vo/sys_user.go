package vo

import (
	"github.com/yx1126/go-admin/app/model"
)

type UserParam struct {
	DeptId   string `json:"deptId" form:"deptId"`
	DeptName string `json:"deptName" form:"deptName"`
	UserName string `json:"userName" form:"userName"`
	NickName string `json:"nickName" form:"nickName"`
	Status   string `json:"status" form:"status"`
}

type UserPagingParam struct {
	UserParam
	PagingVo
}

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

type baseCUVo struct {
	DeptId   *int   `json:"deptId"`
	DeptName string `json:"deptName"`
	NickName string `json:"nickName"`
	UserType string `json:"userType"`
	Email    string `json:"email" binding:"omitempty,email"`
	Phone    string `json:"phone"`
	Sex      string `json:"sex"`
	Avatar   string `json:"avatar"`
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

type UpdatePwdVo struct {
	Id          int    `json:"id"`
	OldPassword string `json:"oldPassword" binding:"required"`
	Password    string `json:"password" binding:"required"`
	ConfirmPwd  string `json:"confirmPwd" binding:"required,eqfield=Password"`
}
