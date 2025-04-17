package systemcontroller

import (
	"github.com/gin-gonic/gin"
	"github.com/yx1126/go-admin/app/service"
	"github.com/yx1126/go-admin/app/vo"
	"github.com/yx1126/go-admin/response"
)

type MenuController struct{}

func (*MenuController) QueryTreeList(c *gin.Context) {
	response.New((&service.MenuService{}).QueryMenuTree(vo.MenuQueryVo{
		Title:  c.Query("title"),
		Status: c.Query("status"),
	})).Json(c)
}

func (*MenuController) QuerySelectTreeList(c *gin.Context) {
	response.New((&service.MenuService{}).QueryMenuSelectTree()).Json(c)
}

func (*MenuController) Create(c *gin.Context) {
	var menu vo.CreateMenuVo
	err := c.ShouldBindJSON(&menu)
	if err != nil {
		response.NewError(err.Error()).Json(c)
		return
	}
	if (&service.MenuService{}).MenuHasSameName(menu.Name, nil) {
		response.NewError(nil).SetMsg("组件名称已存在").Json(c)
		return
	}
	response.New(nil, (&service.MenuService{}).CreateMenu(menu)).Json(c)
}

func (*MenuController) Update(c *gin.Context) {
	var menu vo.UpdateMenuVo
	err := c.ShouldBindJSON(&menu)
	if err != nil {
		response.NewError(err.Error()).Json(c)
		return
	}
	menuId := int(menu.Id)
	if (&service.MenuService{}).MenuHasSameName(menu.Name, &menuId) {
		response.NewError(nil).SetMsg("组件名称已存在").Json(c)
		return
	}
	response.New(nil, (&service.MenuService{}).UpdateMenu(menu)).Json(c)
}

func (*MenuController) Delete(c *gin.Context) {
	var ids []int
	err := c.ShouldBindJSON(&ids)
	if err != nil {
		response.NewError(err.Error()).Json(c)
		return
	}
	for _, id := range ids {
		if (&service.MenuService{}).MenuHasChildren(id) {
			response.NewError(nil).SetMsg("存在子菜单，不允许删除").Json(c)
			return
		}
	}
	response.New(nil, (&service.MenuService{}).DeleteMenus(ids)).Json(c)
}
