package redis

import "github.com/yx1126/go-admin/config"

var (
	// 验证码
	CaptchaCodeKey = config.System.Name + ":captcha:code:"
	// 登录账户密码错误次数
	LoginPwdErrorKey = config.System.Name + ":login:password:error:"
	// 登录用户
	UserTokenKey = config.System.Name + ":user:login:token:"
	// 字典表数据
	SysDictKey = config.System.Name + ":system:dict:data"
)
