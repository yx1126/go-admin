package vo

import "github.com/yx1126/go-admin/app/model"

type RoleParam struct {
	Name   string `json:"name" form:"name"`
	Key    string `json:"key" form:"key"`
	Status string `json:"status" form:"status"`
	PagingVo
}

type RoleVo struct {
	Name   string `json:"name"`
	Key    string `json:"key"`
	Sort   int    `json:"sort"`
	Status string `json:"status"`
	Remark string `json:"remark"`
	model.BaseModel
}

type CreateRoleVo struct {
	Name   string `json:"name"`
	Key    string `json:"key"`
	Sort   int    `json:"sort"`
	Status string `json:"status"`
	Remark string `json:"remark"`
}

type UpdateRoleVo struct {
	Id int `json:"id"`
	CreateRoleVo
}
