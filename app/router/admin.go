package router

import (
	"go-admin/app/controller"
	filecontroller "go-admin/app/controller/file"
	sys "go-admin/app/controller/system"
	mw "go-admin/app/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterAdminRouters(r *gin.RouterGroup) {
	authCtrl := controller.AuthController{}
	r.GET("/code", authCtrl.Code)
	r.POST("/login", authCtrl.Login)
	r.POST("/logout", authCtrl.Logout)
	r.GET("/file/getAvatar", (&filecontroller.FileController{}).GetFileObject)
	r.Use(mw.AuthMiddleware())
	// upload
	file := r.Group("/file")
	{
		file.POST("/uploadAvatar", (&filecontroller.FileController{}).UploadAvatar)
	}

	// auth
	auth := r.Group("/auth")
	{
		auth.GET("/getUserInfo", (&sys.UserController{}).QueryUserInfo)
		auth.GET("/permission", authCtrl.QueryPermission)
	}
	// 系统设置模块
	system := r.Group("/system")
	{
		// 用户管理
		userCtrl := sys.UserController{}
		system.GET("/user", mw.HasPerm("system:user:list"), userCtrl.QueryUserList)
		system.GET("/user/all", mw.HasPerm("system:user:list"), userCtrl.QueryUserAllList)
		system.GET("/user/:id", mw.HasPerm("system:user:list"), userCtrl.QueryUserInfoById)
		system.POST("/user", mw.HasPerm("system:user:add"), userCtrl.Create)
		system.PUT("/user", mw.HasPerm("system:user:edit"), userCtrl.Update)
		system.PUT("/user/reset/:id", mw.HasPerm("system:user:reset"), userCtrl.ResetPwd)
		system.DELETE("/user", mw.HasPerm("system:user:del"), userCtrl.Delete)
		// 角色管理
		roleCtrl := sys.RoleController{}
		system.GET("/role", mw.HasPerm("system:role:list"), roleCtrl.QueryRoleList)
		system.GET("/role/all", mw.HasPerm("system:role:list"), roleCtrl.QueryRoleAllList)
		system.GET("/role/:id", mw.HasPerm("system:role:list"), roleCtrl.QueryRoleInfoById)
		system.POST("/role", mw.HasPerm("system:role:add"), roleCtrl.Create)
		system.PUT("/role", mw.HasPerm("system:role:edit"), roleCtrl.Update)
		system.DELETE("/role", mw.HasPerm("system:role:del"), roleCtrl.Delete)
		// 菜单管理
		menuCtrl := sys.MenuController{}
		system.GET("/menu", mw.HasPerm("system:menu:list"), menuCtrl.QueryTree)
		system.GET("/menu/all", mw.HasPerm("system:menu:list"), menuCtrl.QueryAllTree)
		system.GET("/menu/selectTree", mw.HasPerm("system:menu:list"), menuCtrl.QuerySelectTree)
		system.POST("/menu", mw.HasPerm("system:menu:add"), menuCtrl.Create)
		system.PUT("/menu", mw.HasPerm("system:menu:edit"), menuCtrl.Update)
		system.DELETE("/menu", mw.HasPerm("system:menu:del"), menuCtrl.Delete)
		// 部门管理
		deptCtrl := sys.DeptController{}
		system.GET("/dept", mw.HasPerm("system:dept:list"), deptCtrl.QueryTree)
		system.GET("/dept/all", mw.HasPerm("system:dept:list"), deptCtrl.QuerySelectAllTree)
		system.GET("/dept/selectTree", mw.HasPerm("system:dept:list"), deptCtrl.QuerySelectTree)
		system.POST("/dept", mw.HasPerm("system:dept:add"), deptCtrl.Create)
		system.PUT("/dept", mw.HasPerm("system:dept:edit"), deptCtrl.Update)
		system.DELETE("/dept", mw.HasPerm("system:dept:del"), deptCtrl.Delete)
		// 岗位管理
		postCtrl := sys.PostController{}
		system.GET("/post", mw.HasPerm("system:post:list"), postCtrl.QueryPostList)
		system.GET("/post/all", mw.HasPerm("system:post:list"), postCtrl.QueryPostAllList)
		system.POST("/post", mw.HasPerm("system:post:add"), postCtrl.Create)
		system.PUT("/post", mw.HasPerm("system:post:edit"), postCtrl.Update)
		system.DELETE("/post", mw.HasPerm("system:post:del"), postCtrl.Delete)
		// 字典类型
		dictTypeCtrl := sys.DictController{}
		system.GET("/dict", mw.HasPerm("system:dict:list"), dictTypeCtrl.QueryAllList)
		system.POST("/dict", mw.HasPerm("system:dict:add"), dictTypeCtrl.Create)
		system.PUT("/dict", mw.HasPerm("system:dict:edit"), dictTypeCtrl.Update)
		system.DELETE("/dict", mw.HasPerm("system:dict:del"), dictTypeCtrl.Delete)
		// 字典数据
		dict := system.Group("/dict")
		dict.GET("/data", mw.HasPerm("system:dict:data:list"), dictTypeCtrl.QueryDictDataList)
		dict.GET("/data/:dictType", dictTypeCtrl.QueryDictDataListByType)
		dict.POST("/data", mw.HasPerm("system:dict:data:add"), dictTypeCtrl.CreateData)
		dict.PUT("/data", mw.HasPerm("system:dict:data:edit"), dictTypeCtrl.UpdateData)
		dict.DELETE("/data", mw.HasPerm("system:dict:data:del"), dictTypeCtrl.DeleteData)
	}
}
