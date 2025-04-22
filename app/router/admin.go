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
			menu := system.Group("/menu")
			menuCtrl := systemCtrl.MenuController{}
			menu.GET("", menuCtrl.QueryTree)
			menu.GET("/selectTree", menuCtrl.QuerySelectTree)
			menu.POST("", menuCtrl.Create)
			menu.PUT("", menuCtrl.Update)
			menu.DELETE("", menuCtrl.Delete)
		}
		// 字典类型
		{
			dict := system.Group("/dict")
			dictTypeCtrl := systemCtrl.DictController{}
			dict.GET("", dictTypeCtrl.QueryAllList)
			dict.POST("", dictTypeCtrl.Create)
			dict.PUT("", dictTypeCtrl.Update)
			dict.DELETE("", dictTypeCtrl.Delete)
			{
				// 字典数据
				dictData := dict.Group("/data")
				dictData.GET("", dictTypeCtrl.QueryDictDataList)
				dictData.POST("", dictTypeCtrl.CreateData)
				dictData.PUT("", dictTypeCtrl.UpdateData)
				dictData.DELETE("", dictTypeCtrl.DeleteData)
			}
		}
	}
}
