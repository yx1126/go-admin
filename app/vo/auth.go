package vo

import (
	"github.com/go-playground/validator/v10"
)

type LoginVo struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Code     string `json:"code" binding:"required"`
	Uuid     string `json:"uuid" binding:"required"`
}

func (l *LoginVo) Validate() error {
	v := validator.New(validator.WithRequiredStructEnabled())
	return v.Struct(l)
}

type RegisterVo struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	ConfirmPwd string `json:"confirmPwd"`
	Code       string `json:"code"`
	Uuid       string `json:"uuid"`
}
