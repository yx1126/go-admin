package bind

import (
	"reflect"

	"github.com/gin-gonic/gin"
)

func PagingBind[T any](c *gin.Context, param *T) error {
	if err := c.ShouldBindQuery(&param); err != nil {
		return err
	}
	v := reflect.ValueOf(param).Elem()
	pageField := v.FieldByName("PagingVo").FieldByName("Page")
	sizeField := v.FieldByName("PagingVo").FieldByName("Size")
	if pageField.IsValid() && pageField.Kind() == reflect.Int && pageField.Int() <= 0 {
		pageField.SetInt(1)
	}
	if sizeField.IsValid() && pageField.Kind() == reflect.Int && sizeField.Int() <= 0 {
		sizeField.SetInt(10)
	}
	return nil
}
