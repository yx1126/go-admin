package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	config "github.com/yx1126/go-admin/configs"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run(":" + strconv.Itoa(config.Config.Server.Port))
}
