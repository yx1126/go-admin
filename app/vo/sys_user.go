package vo

import (
	"github.com/yx1126/go-admin/app/model"
)

type baseVo struct {
	DeptId   int    `json:"deptId"`
	DeptName string `json:"deptName"`
	NickName string `json:"nickName"`
	UserType string `json:"userType"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Sex      string `json:"sex"`
	Avatar   string `json:"avatar"`
	Status   string `json:"status"`
	Remark   string `json:"remark"`
}

type UserQueryPageVo struct {
	DeptId   string `json:"deptId" form:"deptId"`
	DeptName string `json:"deptName" form:"deptName"`
	UserName string `json:"userName" form:"userName"`
	NickName string `json:"nickName" form:"nickName"`
	Status   string `json:"status" form:"status"`
	PagingVo
}

type UserVo struct {
	Id        int    `json:"id"`
	UserName  string `json:"userName"`
	LoginIp   string `json:"loginIp"`
	LoginDate string `json:"loginDate"`
	DeptId    int    `json:"deptId"`
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

type CreateUserVo struct {
	UserName string `json:"userName"`
	baseVo
}

type UpdateUserVo struct {
	Id int `json:"id"`
	baseVo
}
