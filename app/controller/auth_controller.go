package controller

import (
	"io"

	"github.com/gin-gonic/gin"
	"github.com/yx1126/go-admin/app/vo"
	"github.com/yx1126/go-admin/common/captcha"
	"github.com/yx1126/go-admin/common/crypto"
	"github.com/yx1126/go-admin/response"
)

type AuthController struct{}

func (*AuthController) Code(c *gin.Context) {
	captcha := captcha.NewCaptcha()
	id, image := captcha.Generate()
	response.NewSuccess(gin.H{
		"uuid":  id,
		"image": image,
	}).Json(c)

}

func (*AuthController) Login(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		response.NewError(err).Json(c)
		return
	}
	var loginInfo vo.LoginVo
	if err := crypto.Unmarshal(body, &loginInfo); err != nil {
		response.NewError(err).Json(c)
		return
	}
	if err := loginInfo.Validate(); err != nil {
		response.NewError(err).Json(c)
		return
	}
	response.NewSuccess(loginInfo).Json(c)
}
