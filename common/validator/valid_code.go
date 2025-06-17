package validator

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func ValidCode(v validator.FieldLevel) bool {
	value, ok := v.Field().Interface().(string)
	if !ok {
		return false
	}
	match, _ := regexp.MatchString(`^[a-zA-Z0-9_-]+$`, value)
	return match
}
