package systemcontroller

import (
	"github.com/gin-gonic/gin"
	systemservice "github.com/yx1126/go-admin/app/service/system"
	"github.com/yx1126/go-admin/app/util"
	"github.com/yx1126/go-admin/app/vo"
	"github.com/yx1126/go-admin/common/constant"
	bind "github.com/yx1126/go-admin/common/should_bind"
	"github.com/yx1126/go-admin/response"
)

type DeptController struct{}

// 查询部门列表
func (*DeptController) QueryTree(c *gin.Context) {
	var param vo.DeptParam
	if err := bind.ShouldBindQuery(c, &param); err != nil {
		response.NewError(nil).Json(c)
		return
	}
	deptList, err := (&systemservice.DeptService{}).QueryDeptList(param)
	response.New(util.ListToTree(deptList, 0), err).Json(c)
}

// 查询所有部门树
func (*DeptController) QuerySelectTree(c *gin.Context) {
	response.New((&systemservice.DeptService{}).QueryDeptSelectTree("")).Json(c)
}

// 查询未禁用的部门树
func (*DeptController) QuerySelectAllTree(c *gin.Context) {
	response.New((&systemservice.DeptService{}).QueryDeptSelectTree(constant.STATUS)).Json(c)
}

// 创建部门
func (*DeptController) Create(c *gin.Context) {
	var dept vo.CreateDeptVo
	if err := bind.ShouldBindJSON(c, &dept); err != nil {
		response.NewError(err).Json(c)
		return
	}
	if (&systemservice.DeptService{}).DeptHasSameName(dept.Name, nil) {
		response.NewError(nil).SetMsg("部门名称已存在").Json(c)
		return
	}
	response.New(nil, (&systemservice.DeptService{}).CreateDept(dept)).Json(c)
}

// 更新部门
func (*DeptController) Update(c *gin.Context) {
	var dept vo.UpdateDeptVo
	if err := bind.ShouldBindJSON(c, &dept); err != nil {
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

// 删除部门
func (*DeptController) Delete(c *gin.Context) {
	var ids []int
	if err := bind.BindIds(c, &ids); err != nil {
		response.NewError(err).Json(c)
		return
	}
	response.New(nil, (&systemservice.DeptService{}).DeleteDept(ids)).Json(c)
}
