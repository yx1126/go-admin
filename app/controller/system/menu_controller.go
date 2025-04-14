package systemcontroller

import (
	"github.com/gin-gonic/gin"
	"github.com/yx1126/go-admin/app/service"
	"github.com/yx1126/go-admin/app/vo"
	"github.com/yx1126/go-admin/response"
)

type MenuController struct{}

func (*MenuController) Create(c *gin.Context) {
	var menu vo.CreateMenuVo
	err := c.ShouldBindJSON(&menu)
	if err != nil {
		response.NewError(err.Error()).Json(c)
		return
	}
	err = (&service.MenuService{}).CreateMenu(menu)
	if err != nil {
		response.NewError(err.Error()).Json(c)
		return
	}
	response.NewSuccess(nil).Json(c)
}

func (*MenuController) QueryTreeList(c *gin.Context) {
	menuList, err := (&service.MenuService{}).QueryMenuTree(vo.MenuQueryVo{
		Title:  c.Query("title"),
		Status: c.Query("status"),
	})
	if err != nil {
		response.NewError(err.Error()).Json(c)
		return
	}
	response.NewSuccess(menuList).Json(c)
}

func (*MenuController) QuerySelectTreeList(c *gin.Context) {
	menuList, err := (&service.MenuService{}).QueryMenuSelectTree()
	if err != nil {
		response.NewError(err.Error()).Json(c)
		return
	}
	response.NewSuccess(menuList).Json(c)
}
