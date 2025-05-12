package validator

import (
	"fmt"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	validatortrans "github.com/yx1126/go-admin/common/validator_trans"
)

// https://www.liwenzhou.com/posts/Go/validator-usages/
func RegisterValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// is_code 注册自定义校验
		v.RegisterValidation("is_code", ValidCode)
		// 注册自定义校验信息
		validatortrans.RegisterMessage("is_code", "请输入数字字母_-")
	}
	if err := validatortrans.InitTrans("zh"); err != nil {
		fmt.Printf("init trans failed, err:%v\n", err)
		return
	}
}
