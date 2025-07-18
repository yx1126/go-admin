package bind

import (
	"errors"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yx1126/go-admin/common/validator"
)

// 绑定分页参数
func BindPaging[T any](c *gin.Context, param *T) error {
	if err := c.ShouldBindQuery(&param); err != nil {
		return validator.TransErr(err)
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

func ShouldBindQuery[T any](c *gin.Context, param *T) error {
	if err := c.ShouldBindQuery(&param); err != nil {
		return validator.TransErr(err)
	}
	return nil
}

func ShouldBind[T any](c *gin.Context, param *T) error {
	if err := c.ShouldBind(&param); err != nil {
		return validator.TransErr(err)
	}
	return nil
}

// create/update 使用
func ShouldBindJSON[T any](c *gin.Context, param *T) error {
	if err := c.ShouldBindJSON(&param); err != nil {
		return validator.TransErr(err)
	}
	username := c.GetString("username")
	val := reflect.ValueOf(param).Elem()
	// 设置 创建人
	if field := val.FieldByName("CreatedBy"); field.IsValid() && field.CanSet() && field.Kind() == reflect.String {
		field.SetString(username)
	}
	// 设置 更新人
	if field := val.FieldByName("UpdatedBy"); field.IsValid() && field.CanSet() && field.Kind() == reflect.String {
		field.SetString(username)
	}
	return nil
}

// 绑定删除id
func BindIds(c *gin.Context, ids *[]int) error {
	var strIds []string
	if err := c.ShouldBind(&strIds); err != nil {
		return err
	}
	if len(strIds) <= 0 {
		return errors.New("请选择要删除的数据")
	}
	for _, v := range strIds {
		id, err := strconv.Atoi(v)
		if err != nil {
			return err
		}
		*ids = append(*ids, id)
	}
	return nil
}
