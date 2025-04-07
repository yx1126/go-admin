package vo

import (
	"github.com/yx1126/go-admin/app/model"
	"github.com/yx1126/go-admin/app/util/datetime"
)

type CreateMenuVo struct {
	ParentId   *uint  `json:"parentId"`
	Name       string `json:"name" binding:"required"`
	Type       string `json:"type" binding:"required_with=0,omitempty,gte=0,lte=3"`
	Link       string `json:"link" binding:"omitempty,url"`
	Title      string `json:"title" binding:"required"`
	IsCache    string `json:"isCache"`
	Icon       string `json:"icon"`
	Path       string `json:"path" binding:"required"`
	IsIframe   string `json:"isIframe"`
	Component  string `json:"component"`
	Permission string `json:"permission"`
	Sort       *int   `json:"sort"`
	Visible    string `json:"visible"`
	Status     string `json:"status"`
	model.BaseModel
}

type MenuVo struct {
	ID         uint              `json:"id"`
	ParentId   *uint             `json:"parentId"`
	Name       string            `json:"name"`
	Type       string            `json:"type"`
	Link       string            `json:"link"`
	Title      string            `json:"title"`
	IsCache    string            `json:"isCache"`
	Icon       string            `json:"icon"`
	Path       string            `json:"path"`
	IsIframe   string            `json:"isIframe"`
	Component  string            `json:"component"`
	Permission string            `json:"permission"`
	Sort       *int              `json:"sort"`
	Visible    string            `json:"visible"`
	Status     string            `json:"status"`
	CreatedBy  string            `json:"createBy"`
	CreatedAt  datetime.Datetime `json:"createdAt"`
	UpdatedBy  string            `json:"updateBy"`
	UpdatedAt  datetime.Datetime `json:"updatedAt"`
}

type MenuTreeVo struct {
	MenuVo
	Children []MenuTreeVo `json:"children" gorm:"-"`
}

func (m MenuTreeVo) GetID() int {
	return int(m.ID)
}

func (m MenuTreeVo) GetParentID() int {
	if m.ParentId == nil {
		return 0
	}
	parentID := int(*m.ParentId)
	return parentID
}

func (m MenuTreeVo) SetChildren(children []MenuTreeVo) MenuTreeVo {
	m.Children = children
	return m
}

type MenuQueryVo struct {
	Title  string `json:"title"`
	Status string `json:"status"`
}
