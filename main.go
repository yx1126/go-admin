package main

import (
	"strconv"

	"go-admin/DB"
	"go-admin/app/router"
	"go-admin/common/validator"
	"go-admin/config"

	"github.com/gin-gonic/gin"
)

func main() {
	// DB
	DB.InitGorm()
	DB.InitRedis()
	DB.InitMinio()
	// validator
	validator.RegisterValidator()
	// mode
	gin.SetMode(config.Server.Mode)
	r := gin.Default()
	// router
	router.Register(r)
	// run
	r.Run(":" + strconv.Itoa(config.Server.Port))
}
