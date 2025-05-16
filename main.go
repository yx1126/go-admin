package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yx1126/go-admin/DB"
	"github.com/yx1126/go-admin/app/router"
	"github.com/yx1126/go-admin/common/validator"
	"github.com/yx1126/go-admin/config"
)

func main() {
	// DB
	DB.InitGorm()
	DB.InitRedis()
	// mode
	gin.SetMode(config.Server.Mode)
	r := gin.Default()
	// validator
	validator.RegisterValidator()
	// router
	router.Register(r)
	// run
	r.Run(":" + strconv.Itoa(config.Server.Port))
}
