package systemcontroller

import (
	"github.com/gin-gonic/gin"
	systemservice "github.com/yx1126/go-admin/app/service/system"
	"github.com/yx1126/go-admin/app/util"
	"github.com/yx1126/go-admin/app/vo"
	"github.com/yx1126/go-admin/response"
)

type DeptController struct{}

// 查询部门列表
func (*DeptController) QueryTree(c *gin.Context) {
	var param vo.DeptParam
	if err := c.ShouldBindQuery(&param); err != nil {
		response.NewError(nil).Json(c)
		return
	}
	deptList, err := (&systemservice.DeptService{}).QueryDeptList(param)
	response.New(util.ListToTree(deptList, 0), err).Json(c)
}

func (*DeptController) QuerySelectTree(c *gin.Context) {
	response.New((&systemservice.DeptService{}).QueryDeptSelectTree()).Json(c)
}

func (*DeptController) Create(c *gin.Context) {
	var dept vo.CreateDeptVo
	if err := c.ShouldBindJSON(&dept); err != nil {
		response.NewError(err).Json(c)
		return
	}
	if (&systemservice.DeptService{}).DeptHasSameName(dept.Name, nil) {
		response.NewError(nil).SetMsg("部门名称已存在").Json(c)
		return
	}
	response.New(nil, (&systemservice.DeptService{}).CreateDept(dept)).Json(c)
}

func (*DeptController) Update(c *gin.Context) {
	var dept vo.UpdateDeptVo
	if err := c.ShouldBindJSON(&dept); err != nil {
		response.NewError(err).Json(c)
		return
	}
	if dept.Id == dept.ParentId {
		response.NewError(nil).SetMsg("请选择正确的父级部门").Json(c)
		return
	}
	if (&systemservice.DeptService{}).DeptHasSameName(dept.Name, &dept.Id) {
		response.NewError(nil).SetMsg("部门名称已存在").Json(c)
		return
	}
	response.New(nil, (&systemservice.DeptService{}).UpdateDept(dept)).Json(c)
}

func (*DeptController) Delete(c *gin.Context) {
	var ids []int
	if err := c.ShouldBindJSON(&ids); err != nil {
		response.NewError(err).Json(c)
		return
	}
	if len(ids) == 0 {
		response.NewError(nil).SetMsg("请选择要删除的数据").Json(c)
		return
	}
	response.New(nil, (&systemservice.DeptService{}).DeleteDept(ids)).Json(c)
}
