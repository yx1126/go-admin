package controller

import (
	"encoding/base64"
	"encoding/json"
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
	cipherData, err := base64.StdEncoding.DecodeString(string(body))
	if err != nil {
		response.NewError(err).Json(c)
		return
	}
	bytes, err := crypto.Parse(cipherData)
	if err != nil {
		response.NewError(err).Json(c)
		return
	}
	var loginInfo vo.LoginVo
	if err := json.Unmarshal(bytes, &loginInfo); err != nil {
		response.NewError(err).Json(c)
		return
	}
	response.NewSuccess(loginInfo).Json(c)
}
