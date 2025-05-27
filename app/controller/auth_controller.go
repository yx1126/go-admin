package controller

import (
	"io"

	"github.com/gin-gonic/gin"
	"github.com/yx1126/go-admin/app/vo"
	"github.com/yx1126/go-admin/common/crypto"
	"github.com/yx1126/go-admin/response"
)

type AuthController struct{}

func (*AuthController) Login(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		response.NewError(err).Json(c)
		return
	}
	var loginInfo vo.LoginVo
	if err := crypto.BindParse(body, &loginInfo); err != nil {
		response.NewError(err).Json(c)
		return
	}
	response.NewSuccess(loginInfo).Json(c)
}
