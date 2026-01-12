package vo

import (
	"errors"

	"go-admin/app/model"
)

type MenuParam struct {
	Title  string `json:"title"`
	Status string `json:"status"`
}

type MenuVo struct {
	ParentId   int    `json:"parentId,string"`
	Name       string `json:"name"`
	Type       uint   `json:"type"`
	Link       string `json:"link"`
	Title      string `json:"title"`
	IsCache    string `json:"isCache"`
	Icon       string `json:"icon"`
	Path       string `json:"path"`
	IsIframe   string `json:"isIframe"`
	Component  string `json:"component"`
	Permission string `json:"permission"`
	Sort       int    `json:"sort"`
	Visible    string `json:"visible"`
	Status     string `json:"status"`
	model.BaseModel
}

type MenuTreeVo struct {
	MenuVo
	Children []MenuTreeVo `json:"children" gorm:"-"`
}

func (m MenuTreeVo) GetID() int {
	return m.Id
}

func (m MenuTreeVo) GetParentID() int {
	return m.ParentId
}

func (m MenuTreeVo) SetChildren(children []MenuTreeVo) MenuTreeVo {
	m.Children = children
	return m
}

type CreateMenuVo struct {
	ParentId   int    `json:"parentId,string"`
	Name       string `json:"name"`
	Type       uint   `json:"type" binding:"required_with=0,omitempty,gte=0,lte=3"`
	Link       string `json:"link" binding:"omitempty,url"`
	Title      string `json:"title" binding:"required"`
	IsCache    string `json:"isCache"`
	Icon       string `json:"icon"`
	Path       string `json:"path"`
	IsIframe   string `json:"isIframe"`
	Component  string `json:"component"`
	Permission string `json:"permission"`
	Sort       int    `json:"sort"`
	Visible    string `json:"visible"`
	Status     string `json:"status"`
	CreatedBy  string
}

func (c *CreateMenuVo) Valid() error {
	if (c.Type == 0 || c.Type == 1) && c.Name == "" {
		return errors.New("组件名称不能为空")
	}
	if c.Type != 3 && c.Path == "" {
		return errors.New("组件路径不能为空")
	}
	return nil
}

type UpdateMenuVo struct {
	BaseVo
	CreateMenuVo
	UpdatedBy string
}
