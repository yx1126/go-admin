package vo

import "github.com/yx1126/go-admin/app/model"

type PostPagingParam struct {
	Name   string `json:"name" form:"name"`
	Code   string `json:"code" form:"code"`
	Status string `json:"status" form:"status"`
	PagingVo
}

type PostVo struct {
	Code   string `json:"code"`
	Name   string `json:"name"`
	Sort   int    `json:"sort"`
	Status string `json:"status"`
	Remark string `json:"remark"`
	model.BaseModel
}

type CreatePostVo struct {
	Code   string `json:"code" binding:"required,is_code"`
	Name   string `json:"name" binding:"required"`
	Sort   int    `json:"sort"`
	Status string `json:"status"`
	Remark string `json:"remark"`
}

type UpdatePostVo struct {
	Id int `json:"id" binding:"required"`
	CreatePostVo
}
