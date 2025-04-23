package vo

import "github.com/yx1126/go-admin/app/model"

type DictTypeListVo struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	NodeType string `json:"nodeType"`
	Status   string `json:"status"`
	Remark   string `json:"remark"`
	model.BaseModel
}

type CreateDictType struct {
	Name     string `json:"name" binding:"required"`
	Type     string `json:"type" binding:"required"`
	NodeType string `json:"nodeType"`
	Status   string `json:"status"`
	Remark   string `json:"remark"`
}

type UpdateDictType struct {
	Id uint `json:"id"`
	CreateDictType
}

type DictDataListVo struct {
	Id        uint   `json:"id"`
	DictId    uint   `json:"dictId"`
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
	DictId    uint   `json:"dictId" binding:"required"`
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
	Id uint `json:"id"`
	CreateDictData
}
