package systemcontroller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	systemservice "github.com/yx1126/go-admin/app/service/system"
	"github.com/yx1126/go-admin/app/vo"
	"github.com/yx1126/go-admin/common/password"
	bind "github.com/yx1126/go-admin/common/should_bind"
	"github.com/yx1126/go-admin/common/util"
	"github.com/yx1126/go-admin/config"
	"github.com/yx1126/go-admin/response"
)

type UserController struct{}

// 获取用户列表
func (*UserController) QueryUserList(c *gin.Context) {
	var params vo.UserPagingParam
	if err := bind.BindPaging(c, &params); err != nil {
		response.NewError(err).Json(c)
		return
	}
	data, err := (&systemservice.UserService{}).QueryUserList(params)
	paging := response.Paging{
		List:  data.Data,
		Page:  params.Page,
		Size:  params.Size,
		Total: data.Count,
	}
	response.New(nil, err).SetPaging(paging).Json(c)
}

// 查询所有用户列表
func (*UserController) QueryUserAllList(c *gin.Context) {
	var params vo.UserParam
	if err := bind.ShouldBindQuery(c, &params); err != nil {
		response.NewError(err).Json(c)
		return
	}
	response.New((&systemservice.UserService{}).QueryUserAllList(params)).Json(c)
}

// 根据id获取用户信息
func (*UserController) QueryUserInfoById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.NewError(err).Json(c)
		return
	}
	response.New((&systemservice.UserService{}).QueryUserById(id)).Json(c)
}

// 新增用户
func (*UserController) Create(c *gin.Context) {
	var user vo.CreateUserVo
	if err := bind.ShouldBindJSON(c, &user); err != nil {
		response.NewError(err).Json(c)
		return
	}
	if ok := (&systemservice.UserService{}).IsHasSameName(user.UserName); ok {
		response.NewError(nil).SetMsg("用户名已存在").Json(c)
		return
	}
	response.New(nil, (&systemservice.UserService{}).CreateUser(user)).Json(c)
}

// 更新用户
func (*UserController) Update(c *gin.Context) {
	var user vo.UpdateUserVo
	if err := bind.ShouldBindJSON(c, &user); err != nil {
		response.NewError(err).Json(c)
		return
	}
	response.New(nil, (&systemservice.UserService{}).UpdateUser(user)).Json(c)
}

// 删除用户
func (*UserController) Delete(c *gin.Context) {
	var ids []int
	if err := bind.BindIds(c, &ids); err != nil {
		response.NewError(err).Json(c)
		return
	}
	if util.Contains(ids, 1) {
		response.NewError(nil).SetMsg("删除失败，禁止删除超超级管理员！").Json(c)
		return
	}
	response.New(nil, (&systemservice.UserService{}).DeleteUser(ids)).Json(c)
}

// 重置密码
func (*UserController) ResetPwd(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.NewError(err).Json(c)
		return
	}
	pwd, _ := password.Encode(config.User.Password)
	response.New(nil, (&systemservice.UserService{}).UpdatePwd(id, pwd)).Json(c)
}

// 修改密码
func (*UserController) UpdatePwd(c *gin.Context) {
	var params vo.UpdatePwdVo
	if err := bind.ShouldBindJSON(c, &params); err != nil {
		response.NewError(err).Json(c)
		return
	}
	response.New(nil, (&systemservice.UserService{}).UpdatePwd(params.Id, params.Password)).Json(c)
}
