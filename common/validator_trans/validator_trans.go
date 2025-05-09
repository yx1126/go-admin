package validatortrans

import (
	"fmt"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)

var Trans ut.Translator

func InitTrans(locales string) error {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
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
			return enTranslations.RegisterDefaultTranslations(v, Trans)
		case "zh":
			return zhTranslations.RegisterDefaultTranslations(v, Trans)
		default:
			return enTranslations.RegisterDefaultTranslations(v, Trans)
		}
	}
	return nil
}
