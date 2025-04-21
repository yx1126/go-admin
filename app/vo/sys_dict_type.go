package vo

import "github.com/yx1126/go-admin/app/model"

type DictTypeListVo struct {
	Id     uint   `json:"id"`
	Name   string `json:"name"`
	Type   string `json:"type"`
	Status string `json:"status"`
	Remark string `json:"remark"`
	model.BaseModel
}

type CreateDictType struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Status string `json:"status"`
	Remark string `json:"remark"`
}

type UpdateDictType struct {
	Id uint `json:"id"`
	CreateDictType
}
