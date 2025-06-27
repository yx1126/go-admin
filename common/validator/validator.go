package validator

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

var Validator *validator.Validate

func removeTopStruct(fields map[string]string) map[string]string {
	res := map[string]string{}
	for field, err := range fields {
		res[field[strings.Index(field, ".")+1:]] = err
	}
	return res
}

// validator 翻译
func TransErr(err error) error {
	// 获取validator.ValidationErrors类型的errors
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		// 非validator.ValidationErrors类型错误直接返回
		return err
	}
	// validator.ValidationErrors类型错误则进行翻译
	msg, err := json.Marshal(removeTopStruct(errs.Translate(Trans)))
	if err != nil {
		return err
	}
	return errors.New(string(msg))
}

func Struct(s any) error {
	err := Validator.Struct(s)
	return TransErr(err)
}

// https://www.liwenzhou.com/posts/Go/validator-usages/
func RegisterValidator() {
	// Validator = validator.New()
	// Validator.SetTagName("binding")

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 注册自定义校验信息
		RegisterMessage("is_code", "请输入数字字母_-")
		// is_code 注册自定义校验
		v.RegisterValidation("is_code", ValidCode)
		// gin trans zh
		if err := InitTran(v, "zh"); err != nil {
			fmt.Printf("init trans failed, err:%v\n", err)
			return
		}
		Validator = v
	}
	// local
	// if err := InitTran(Validator, "zh"); err != nil {
	// 	fmt.Printf("init trans failed, err:%v\n", err)
	// 	return
	// }
}
