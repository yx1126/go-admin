package vo

import "github.com/yx1126/go-admin/app/model"

type DeptParam struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}

type DeptVo struct {
	ParentId int      `json:"parentId"`
	Name     string   `json:"name"`
	Sort     int      `json:"sort"`
	LeaderId int      `json:"leaderId"`
	Status   string   `json:"status"`
	Children []DeptVo `json:"children" gorm:"-"`
	model.BaseModel
}

func (d DeptVo) GetID() int {
	return d.Id
}

func (d DeptVo) GetParentID() int {
	return int(d.ParentId)
}

func (d DeptVo) SetChildren(children []DeptVo) DeptVo {
	d.Children = children
	return d
}

type DeptTreeVo struct {
	ParentId int          `json:"parentId"`
	Name     string       `json:"name"`
	Sort     int          `json:"sort"`
	LeaderId *int         `json:"leaderId"`
	Leader   *UserVo      `json:"leader" gorm:"foreignKey:LeaderId;references:Id"`
	Status   string       `json:"status"`
	Children []DeptTreeVo `json:"children" gorm:"-"`
	model.BaseModel
}

func (d DeptTreeVo) GetID() int {
	return d.Id
}

func (d DeptTreeVo) GetParentID() int {
	return int(d.ParentId)
}

func (d DeptTreeVo) SetChildren(children []DeptTreeVo) DeptTreeVo {
	d.Children = children
	return d
}

type CreateDeptVo struct {
	ParentId int    `json:"parentId"`
	Name     string `json:"name" binding:"required"`
	Sort     int    `json:"sort"`
	LeaderId *int   `json:"leaderId" binding:"omitempty"`
	Status   string `json:"status"`
}

type UpdateDeptVo struct {
	Id int `json:"id"`
	CreateDeptVo
}
