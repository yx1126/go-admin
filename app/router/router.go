package router

import "github.com/gin-gonic/gin"

func Register(g *gin.Engine) {
	api := g.Group("/api")
	RegisterAdminRouters(api)
}
