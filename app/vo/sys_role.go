package vo

import (
	"github.com/yx1126/go-admin/app/model"
	"github.com/yx1126/go-admin/common/types"
)

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

type RoleInfoVo struct {
	MenuIds *[]types.Long `json:"menuIds"`
	RoleVo
}

type CreateRoleVo struct {
	Name    string        `json:"name" binding:"required"`
	Key     string        `json:"key" binding:"required,is_code"`
	Sort    int           `json:"sort"`
	Status  string        `json:"status"`
	Remark  string        `json:"remark"`
	MenuIds *[]types.Long `json:"menuIds"`
}

type UpdateRoleVo struct {
	BaseVo
	CreateRoleVo
}
