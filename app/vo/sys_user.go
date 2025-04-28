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
	Password string `json:"password"`
	Status   string `json:"status"`
	Remark   string `json:"remark"`
}

type UserQueryPageVo struct {
	DeptId   string `json:"deptId"`
	UserName string `json:"userName"`
	NickName string `json:"nickName"`
	Status   string `json:"status"`
	PagingVo
}

type UserVo struct {
	Id        int    `json:"id"`
	UserName  string `json:"userName"`
	LoginIp   string `json:"loginIp"`
	LoginDate string `json:"loginDate"`
	baseVo
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
