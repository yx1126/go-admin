package controller

import (
	"fmt"
	"io"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yx1126/go-admin/DB"
	systemservice "github.com/yx1126/go-admin/app/service/system"
	"github.com/yx1126/go-admin/app/vo"
	"github.com/yx1126/go-admin/common/captcha"
	"github.com/yx1126/go-admin/common/constant"
	"github.com/yx1126/go-admin/common/crypto"
	"github.com/yx1126/go-admin/common/password"
	"github.com/yx1126/go-admin/common/redis"
	"github.com/yx1126/go-admin/common/token"
	"github.com/yx1126/go-admin/config"
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
	// 解密
	if err := crypto.Unmarshal(body, &loginInfo); err != nil {
		response.NewError(err).Json(c)
		return
	}
	// 字段校验
	if err := loginInfo.Validate(); err != nil {
		response.NewError(err).Json(c)
		return
	}
	// 验证码校验
	if ok := captcha.NewCaptcha().Verify(loginInfo.Uuid, loginInfo.Code); !ok {
		response.NewError(nil).SetMsg("验证码错误").Json(c)
		return
	}
	user, err := (&systemservice.UserService{}).QueryUserPwdByUsername(loginInfo.Username)
	if err != nil {
		response.NewError(nil).SetMsg("用户名或密码错误").Json(c)
		return
	}
	if user.Status != constant.STATUS {
		response.NewError(nil).SetMsg("用户已被禁用，请联系管理员").Json(c)
		return
	}
	redisKey := redis.LoginPwdErrorKey + loginInfo.Username
	count, _ := DB.Redis.Ctx.Get(c.Request.Context(), redisKey).Int()
	if count >= config.User.MaxRetryCount {
		response.NewError(nil).SetMsg("密码错误次数超过限制，请稍后再试").Json(c)
		return
	}
	if ok := password.Matches(loginInfo.Password, user.Password); !ok {
		DB.Redis.Ctx.Set(c.Request.Context(), redisKey, count+1, time.Minute*time.Duration(config.User.LockTime))
		response.NewError(nil).SetMsg("用户名或密码错误").Json(c)
		return
	}
	// 登录成功，删除错误次数
	DB.Redis.Ctx.Del(c.Request.Context(), redisKey)

	token, err := token.GenToken(user.Id, user.UserName)
	if err != nil {
		fmt.Println("token -->", err.Error())
		response.NewError(err).Json(c)
		return
	}
	// 更新登录的ip和时间
	response.NewSuccess(token).Json(c)
}
