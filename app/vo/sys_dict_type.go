package vo

import "github.com/yx1126/go-admin/app/model"

type DictTypeListVo struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	NodeType string `json:"nodeType"`
	Status   string `json:"status"`
	Remark   string `json:"remark"`
	model.BaseModel
}

type CreateDictType struct {
	Name     string `json:"name" binding:"required"`
	Type     string `json:"type" binding:"required,is_code"`
	NodeType string `json:"nodeType"`
	Status   string `json:"status"`
	Remark   string `json:"remark"`
}

type UpdateDictType struct {
	BaseVo
	CreateDictType
}

type DictPagingParam struct {
	Id     *int   `json:"id" form:"id"`
	Label  string `json:"label" form:"label"`
	Status string `json:"status" form:"status"`
	PagingVo
}

type DictDataListVo struct {
	DictId    int    `json:"dictId"`
	Sort      int    `json:"sort"`
	Label     string `json:"label"`
	Value     string `json:"value"`
	Type      string `json:"type"`
	CssClass  string `json:"cssClass"`
	ListClass string `json:"listClass"`
	IsDefault string `json:"isDefault"`
	Status    string `json:"status"`
	Remark    string `json:"remark"`

	DictType string `json:"dictType"`
	NodeType string `json:"nodeType"`

	model.BaseModel
}

type CreateDictData struct {
	DictId    int    `json:"dictId" binding:"required"`
	Sort      int    `json:"sort"`
	Label     string `json:"label" binding:"required"`
	Value     string `json:"value" binding:"required"`
	Type      string `json:"type"`
	CssClass  string `json:"cssClass"`
	ListClass string `json:"listClass"`
	IsDefault string `json:"isDefault"`
	Status    string `json:"status"`
	Remark    string `json:"remark"`
}

type UpdateDictData struct {
	BaseVo
	CreateDictData
}
