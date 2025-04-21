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
			menu.GET("", menuCtrl.QueryTree)
			menu.GET("/selectTree", menuCtrl.QuerySelectTree)
			menu.POST("", menuCtrl.Create)
			menu.PUT("", menuCtrl.Update)
			menu.DELETE("", menuCtrl.Delete)
		}
		// 字典类型
		{
			dictType := system.Group("dictType")
			dictTypeCtrl := systemCtrl.DictTypeController{}
			dictType.GET("", dictTypeCtrl.QueryAllList)
			dictType.POST("", dictTypeCtrl.Create)
			dictType.PUT("", dictTypeCtrl.Update)
			dictType.DELETE("", dictTypeCtrl.Delete)
		}
	}
}
