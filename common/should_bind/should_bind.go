package bind

import (
	"encoding/json"
	"errors"
	"reflect"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	validatortrans "github.com/yx1126/go-admin/common/validator_trans"
)

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
	msg, err := json.Marshal(removeTopStruct(errs.Translate(validatortrans.Trans)))
	if err != nil {
		return err
	}
	return errors.New(string(msg))
}

// 绑定分页参数
func BindPaging[T any](c *gin.Context, param *T) error {
	if err := c.ShouldBindQuery(&param); err != nil {
		return TransErr(err)
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
		return TransErr(err)
	}
	return nil
}

func ShouldBind[T any](c *gin.Context, param *T) error {
	if err := c.ShouldBind(&param); err != nil {
		return TransErr(err)
	}
	return nil
}

func ShouldBindJSON[T any](c *gin.Context, param *T) error {
	if err := c.ShouldBindJSON(&param); err != nil {
		return TransErr(err)
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
