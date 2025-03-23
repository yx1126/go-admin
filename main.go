package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/yx1126/go-admin/configs"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
