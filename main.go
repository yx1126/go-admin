package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yx1126/go-admin/DB"
	"github.com/yx1126/go-admin/app/router"
	"github.com/yx1126/go-admin/config"
)

func main() {
	DB.InitGorm()
	r := gin.Default()
	router.Register(r)
	r.Run(":" + strconv.Itoa(config.Config.Server.Port))
}
