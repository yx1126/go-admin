package vo

import (
	"time"

	"github.com/yx1126/go-admin/app/model"
)

type CreateMenuVo struct {
	ParentId   *int   `json:"parentId"`
	Name       string `json:"name" binding:"required"`
	Type       int    `json:"type" binding:"required_with=0,omitempty,gte=0,lte=3"`
	Link       string `json:"link" binding:"omitempty,url"`
	Title      string `json:"title" binding:"required"`
	IsCache    int    `json:"isCache"`
	Icon       string `json:"icon"`
	Path       string `json:"path" binding:"required"`
	IsIframe   int    `json:"isIframe"`
	Component  string `json:"component"`
	Permission string `json:"permission"`
	Sort       *int   `json:"sort"`
	Visible    int    `json:"visible"`
	model.BaseModel
}

type MenuVo struct {
	ID         uint      `json:"id"`
	ParentId   int       `json:"parentId"`
	Name       string    `json:"name"`
	Type       int       `json:"type"`
	Link       string    `json:"link"`
	Title      string    `json:"title"`
	IsCache    int       `json:"isCache"`
	Icon       string    `json:"icon"`
	Path       string    `json:"path"`
	IsIframe   int       `json:"isIframe"`
	Component  string    `json:"component"`
	Permission string    `json:"permission"`
	Sort       int       `json:"sort"`
	Visible    int       `json:"visible"`
	CreatedBy  string    `json:"createBy"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedBy  string    `json:"updateBy"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

type MenuTreeVo struct {
	MenuVo
	Children []MenuVo `json:"children"`
}
