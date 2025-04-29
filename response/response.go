package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Paging struct {
	Total int         `json:"total"`
	Page  int         `json:"page"`
	Size  int         `json:"size"`
	Data  interface{} `json:"data"`
}

type Response struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func (r *Response) SetMsg(message string) *Response {
	r.Message = message
	return r
}

func (r *Response) SetCode(code int) *Response {
	r.Code = code
	return r
}

func (r *Response) SetData(data interface{}) *Response {
	r.Data = data
	return r
}

func (r *Response) SetPaging(paging Paging) *Response {
	r.SetData(paging)
	return r
}

func (r *Response) Json(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    r.Code,
		"data":    r.Data,
		"message": r.Message,
	})
}

func New(data interface{}, err error) *Response {
	if err != nil {
		return NewError(err)
	}
	return NewSuccess(data)
}

func NewSuccess(data interface{}) *Response {
	return &Response{
		Code:    200,
		Data:    data,
		Message: "成功！",
	}
}

func NewError(err error) *Response {
	res := &Response{
		Code: 500,
		Data: nil,
	}
	if err != nil {
		res.Message = err.Error()
	} else {
		res.Message = "失败！"
	}
	return res
}
