package main

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yx1126/go-admin/DB"
	"github.com/yx1126/go-admin/app/router"
	"github.com/yx1126/go-admin/common/validator"
	validatortrans "github.com/yx1126/go-admin/common/validator_trans"
	"github.com/yx1126/go-admin/config"
)

func main() {
	if err := validatortrans.InitTrans("zh"); err != nil {
		fmt.Printf("init trans failed, err:%v\n", err)
		return
	}
	DB.InitGorm()
	r := gin.Default()
	validator.RegisterValidator()
	router.Register(r)
	r.Run(":" + strconv.Itoa(config.Server.Port))
}
