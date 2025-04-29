package systemcontroller

import (
	"github.com/gin-gonic/gin"
	service "github.com/yx1126/go-admin/app/service/system"
	"github.com/yx1126/go-admin/app/vo"
	"github.com/yx1126/go-admin/response"
)

type UserController struct{}

func (*UserController) QueryUserList(c *gin.Context) {
	var params vo.UserQueryPageVo
	if err := c.ShouldBindQuery(&params); err != nil {
		response.NewError(err).Json(c)
		return
	}
	userList, total, err := (&service.UserService{}).QueryUserList(params)
	paging := response.Paging{
		Data:  userList,
		Page:  params.Page,
		Size:  params.Size,
		Total: total,
	}
	if err != nil {
		response.NewError(err).SetPaging(paging).Json(c)
		return
	}
	response.NewSuccess(paging).Json(c)
}

func (*UserController) Create(c *gin.Context) {
	var user vo.CreateUserVo
	if err := c.ShouldBindJSON(&user); err != nil {
		response.NewError(err).Json(c)
		return
	}
	response.New(nil, (&service.UserService{}).CreateUser(user)).Json(c)
}

func (*UserController) Update(c *gin.Context) {
	var user vo.UpdateUserVo
	if err := c.ShouldBindJSON(&user); err != nil {
		response.NewError(err).Json(c)
		return
	}
	response.New(nil, (&service.UserService{}).UpdateUser(user)).Json(c)
}

func (*UserController) Delete(c *gin.Context) {
	var ids []int
	if err := c.ShouldBindJSON(&ids); err != nil {
		response.NewError(err).Json(c)
		return
	}
	response.New(nil, (&service.UserService{}).DeleteUser(ids)).Json(c)
}
