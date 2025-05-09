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

type MenuController struct{}

// 菜单树查询
func (*MenuController) QueryTree(c *gin.Context) {
	menuList, err := (&systemservice.MenuService{}).QueryMenuList(vo.MenuParam{
		Title:  c.Query("title"),
		Status: c.Query("status"),
	})
	response.New(util.ListToTree(menuList, 0), err).Json(c)
}

// 菜单页下拉列表查询
func (*MenuController) QuerySelectTree(c *gin.Context) {
	response.New((&systemservice.MenuService{}).QueryMenuSelectTree("")).Json(c)
}

func (*MenuController) QueryAllTree(c *gin.Context) {
	response.New((&systemservice.MenuService{}).QueryMenuSelectTree(constant.STATUS)).Json(c)
}

// 菜单新增
func (*MenuController) Create(c *gin.Context) {
	var menu vo.CreateMenuVo
	err := bind.ShouldBindJSON(c, &menu)
	if err != nil {
		response.NewError(err).Json(c)
		return
	}
	if (&systemservice.MenuService{}).IsHasSameName(menu.Name, nil) {
		response.NewError(nil).SetMsg("组件名称已存在").Json(c)
		return
	}
	response.New(nil, (&systemservice.MenuService{}).CreateMenu(menu)).Json(c)
}

// 菜单更新
func (*MenuController) Update(c *gin.Context) {
	var menu vo.UpdateMenuVo
	err := bind.ShouldBindJSON(c, &menu)
	if err != nil {
		response.NewError(err).Json(c)
		return
	}
	if menu.Id == menu.ParentId {
		response.NewError(nil).SetMsg("请选择正确的父级菜单").Json(c)
		return
	}
	if (&systemservice.MenuService{}).IsHasSameName(menu.Name, &menu.Id) {
		response.NewError(nil).SetMsg("组件名称已存在").Json(c)
		return
	}
	response.New(nil, (&systemservice.MenuService{}).UpdateMenu(menu)).Json(c)
}

// 菜单删除
func (*MenuController) Delete(c *gin.Context) {
	var ids []int
	if err := bind.BindIds(c, &ids); err != nil {
		response.NewError(err).Json(c)
		return
	}
	for _, id := range ids {
		if (&systemservice.MenuService{}).IsHasChildren(id) {
			response.NewError(nil).SetMsg("存在子菜单，不允许删除").Json(c)
			return
		}
	}
	response.New(nil, (&systemservice.MenuService{}).DeleteMenus(ids)).Json(c)
}
