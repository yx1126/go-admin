package bind

import (
	"errors"
	"reflect"

	"github.com/gin-gonic/gin"
)

// 绑定分页参数
func BindPaging[T any](c *gin.Context, param *T) error {
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

// 绑定删除id
func BindIds(c *gin.Context, ids *[]int) error {
	if err := c.ShouldBind(&ids); err != nil {
		return err
	}
	if len(*ids) <= 0 {
		return errors.New("请选择要删除的数据")
	}
	return nil
}
