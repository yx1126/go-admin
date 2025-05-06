package router

import (
	"github.com/gin-gonic/gin"
	systemCtrl "github.com/yx1126/go-admin/app/controller/system"
)

func RegisterAdminRouters(g *gin.RouterGroup) {
	// 系统设置模块
	system := g.Group("/system")
	{
		// 用户管理
		user := system.Group("/user")
		{
			userCtrl := systemCtrl.UserController{}
			user.GET("", userCtrl.QueryUserList)
			user.GET("/all", userCtrl.QueryUserAllList)
			user.GET("/:id", userCtrl.QueryUserInfoById)
			user.POST("", userCtrl.Create)
			user.PUT("", userCtrl.Update)
			user.PUT("/reset/:id", userCtrl.ResetPwd)
			user.DELETE("", userCtrl.Delete)
		}
		// 部门管理
		dept := system.Group("/dept")
		{
			deptCtrl := systemCtrl.DeptController{}
			dept.GET("", deptCtrl.QueryTree)
			dept.GET("/all", deptCtrl.QuerySelectAllTree)
			dept.GET("/selectTree", deptCtrl.QuerySelectTree)
			dept.POST("", deptCtrl.Create)
			dept.PUT("", deptCtrl.Update)
			dept.DELETE("", deptCtrl.Delete)
		}
		// 菜单管理
		menu := system.Group("/menu")
		{
			menuCtrl := systemCtrl.MenuController{}
			menu.GET("", menuCtrl.QueryTree)
			menu.GET("/selectTree", menuCtrl.QuerySelectTree)
			menu.POST("", menuCtrl.Create)
			menu.PUT("", menuCtrl.Update)
			menu.DELETE("", menuCtrl.Delete)
		}
		// 字典类型
		dict := system.Group("/dict")
		{
			dictTypeCtrl := systemCtrl.DictController{}
			dict.GET("", dictTypeCtrl.QueryAllList)
			dict.POST("", dictTypeCtrl.Create)
			dict.PUT("", dictTypeCtrl.Update)
			dict.DELETE("", dictTypeCtrl.Delete)
			// 字典数据
			dictData := dict.Group("/data")
			{
				dictData.GET("", dictTypeCtrl.QueryDictDataList)
				dictData.GET("/:dictType", dictTypeCtrl.QueryDictDataListByType)
				dictData.POST("", dictTypeCtrl.CreateData)
				dictData.PUT("", dictTypeCtrl.UpdateData)
				dictData.DELETE("", dictTypeCtrl.DeleteData)
			}
		}
		// 岗位管理
		post := system.Group("/post")
		{
			postCtrl := systemCtrl.PostController{}
			post.GET("", postCtrl.QueryPostList)
			post.GET("/all", postCtrl.QueryPostAllList)
			post.POST("", postCtrl.Create)
			post.PUT("", postCtrl.Update)
			post.DELETE("", postCtrl.Delete)
		}
	}
}
