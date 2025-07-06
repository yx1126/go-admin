package router

import (
	"github.com/gin-gonic/gin"
	"github.com/yx1126/go-admin/app/controller"
	filecontroller "github.com/yx1126/go-admin/app/controller/file"
	sys "github.com/yx1126/go-admin/app/controller/system"
	mw "github.com/yx1126/go-admin/app/middleware"
)

func RegisterAdminRouters(r *gin.RouterGroup) {
	authCtrl := controller.AuthController{}
	r.GET("/code", authCtrl.Code)
	r.POST("/login", authCtrl.Login)
	r.POST("/logout", authCtrl.Logout)
	r.Use(mw.AuthMiddleware())
	// upload
	file := r.Group("/file")
	{
		file.GET("/getFile", (&filecontroller.FileController{}).GetFileObject)
		file.POST("/uploadAvatar", (&filecontroller.FileController{}).UploadAvatar)
	}

	// auth
	auth := r.Group("/auth")
	{
		auth.GET("/getUserInfo", (&sys.UserController{}).QueryUserInfo)
	}
	// 系统设置模块
	system := r.Group("/system")
	{
		// 用户管理
		userCtrl := sys.UserController{}
		system.GET("/user", userCtrl.QueryUserList)
		system.GET("/user/all", userCtrl.QueryUserAllList)
		system.GET("/user/:id", userCtrl.QueryUserInfoById)
		system.POST("/user", userCtrl.Create)
		system.PUT("/user", userCtrl.Update)
		system.PUT("/user/reset/:id", userCtrl.ResetPwd)
		system.DELETE("/user", userCtrl.Delete)
		// 角色管理
		roleCtrl := sys.RoleController{}
		system.GET("/role", roleCtrl.QueryRoleList)
		system.GET("/role/all", roleCtrl.QueryRoleAllList)
		system.GET("/role/:id", roleCtrl.QueryRoleInfoById)
		system.POST("/role", roleCtrl.Create)
		system.PUT("/role", roleCtrl.Update)
		system.DELETE("/role", roleCtrl.Delete)
		// 菜单管理
		menuCtrl := sys.MenuController{}
		system.GET("/menu", menuCtrl.QueryTree)
		system.GET("/menu/all", menuCtrl.QueryAllTree)
		system.GET("/menu/selectTree", menuCtrl.QuerySelectTree)
		system.POST("/menu", menuCtrl.Create)
		system.PUT("/menu", menuCtrl.Update)
		system.DELETE("/menu", menuCtrl.Delete)
		// 部门管理
		deptCtrl := sys.DeptController{}
		system.GET("/dept", deptCtrl.QueryTree)
		system.GET("/dept/all", deptCtrl.QuerySelectAllTree)
		system.GET("/dept/selectTree", deptCtrl.QuerySelectTree)
		system.POST("/dept", deptCtrl.Create)
		system.PUT("/dept", deptCtrl.Update)
		system.DELETE("/dept", deptCtrl.Delete)
		// 岗位管理
		postCtrl := sys.PostController{}
		system.GET("/post", postCtrl.QueryPostList)
		system.GET("/post/all", postCtrl.QueryPostAllList)
		system.POST("/post", postCtrl.Create)
		system.PUT("/post", postCtrl.Update)
		system.DELETE("/post", postCtrl.Delete)
		// 字典类型
		dictTypeCtrl := sys.DictController{}
		system.GET("/dict", dictTypeCtrl.QueryAllList)
		system.POST("/dict", dictTypeCtrl.Create)
		system.PUT("/dict", dictTypeCtrl.Update)
		system.DELETE("/dict", dictTypeCtrl.Delete)
		// 字典数据
		dict := system.Group("/dict")
		dict.GET("/data", dictTypeCtrl.QueryDictDataList)
		dict.GET("/data/:dictType", dictTypeCtrl.QueryDictDataListByType)
		dict.POST("/data", dictTypeCtrl.CreateData)
		dict.PUT("/data", dictTypeCtrl.UpdateData)
		dict.DELETE("/data", dictTypeCtrl.DeleteData)
	}
}
