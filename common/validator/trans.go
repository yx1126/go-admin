package validator

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)

var Trans ut.Translator

var messages = make(map[string]string, 0)

func translate(trans ut.Translator, fe validator.FieldError) string {
	msg, err := trans.T(fe.Tag(), fe.Field())
	if err != nil {
		panic(fe.(error).Error())
	}
	return msg
}

func registerTranslator(tag string, msg string) validator.RegisterTranslationsFunc {
	return func(trans ut.Translator) error {
		if err := trans.Add(tag, msg, false); err != nil {
			return err
		}
		return nil
	}
}

func RegisterMessage(tag, message string) {
	messages[tag] = message
}

func InitTran(v *validator.Validate, locales string) (err error) {
	// 注册一个获取json tag的自定义方法
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	zhT := zh.New()
	enT := en.New()

	uni := ut.New(enT, zhT, enT)

	var ok bool
	Trans, ok = uni.GetTranslator(locales)
	if !ok {
		return fmt.Errorf("uni.GetTranslator(%s) failed", locales)
	}
	switch locales {
	case "en":
		err = enTranslations.RegisterDefaultTranslations(v, Trans)
	case "zh":
		err = zhTranslations.RegisterDefaultTranslations(v, Trans)
	default:
		err = enTranslations.RegisterDefaultTranslations(v, Trans)
	}
	if err != nil {
		return err
	}
	for tag, message := range messages {
		if err := v.RegisterTranslation(tag, Trans, registerTranslator(tag, message), translate); err != nil {
			return err
		}
	}
	return
}
