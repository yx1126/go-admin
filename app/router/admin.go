package router

import (
	"github.com/gin-gonic/gin"
	systemCtrl "github.com/yx1126/go-admin/app/controller/system"
)

func RegisterAdminRouters(g *gin.RouterGroup) {
	// 系统设置模块
	system := g.Group("/system")
	{
		// 菜单
		{
			menu := system.Group("menu")
			menuCtrl := systemCtrl.MenuController{}
			menu.POST("", menuCtrl.Create)
		}
	}
}
