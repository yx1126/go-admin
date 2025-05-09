package systemcontroller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	systemservice "github.com/yx1126/go-admin/app/service/system"
	"github.com/yx1126/go-admin/app/vo"
	bind "github.com/yx1126/go-admin/common/should_bind"
	"github.com/yx1126/go-admin/response"
)

type RoleController struct{}

// 分页查询
func (*RoleController) QueryRoleList(c *gin.Context) {
	var params vo.RoleParam
	if err := bind.BindPaging(c, &params); err != nil {
		response.NewError(nil).Json(c)
		return
	}
	data, err := (&systemservice.RoleService{}).QueryRoleList(params)
	paging := response.Paging{
		List:  data.Data,
		Page:  params.Page,
		Size:  params.Size,
		Total: data.Count,
	}
	if err != nil {
		response.NewError(err).SetPaging(paging).Json(c)
		return
	}
	response.NewSuccess(paging).Json(c)
}

// 查询全部
func (*RoleController) QueryRoleAllList(c *gin.Context) {
	response.New((&systemservice.RoleService{}).QueryRoleAllList()).Json(c)
}

// 根据id查询角色信息
func (*RoleController) QueryRoleInfoById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.NewError(err).Json(c)
		return
	}
	response.New((&systemservice.RoleService{}).QueryRoleInfo(id)).Json(c)
}

// 创建角色
func (*RoleController) Create(c *gin.Context) {
	var role vo.CreateRoleVo
	if err := c.ShouldBind(&role); err != nil {
		response.NewError(err).Json(c)
		return
	}
	if ok := (&systemservice.RoleService{}).IsHasSameName(role.Name, nil); ok {
		response.NewError(nil).SetMsg("角色名称已存在").Json(c)
		return
	}
	if ok := (&systemservice.RoleService{}).IsHasSameKey(role.Key, nil); ok {
		response.NewError(nil).SetMsg("角色权限字符已存在").Json(c)
		return
	}
	response.New(nil, (&systemservice.RoleService{}).CreateRole(role)).Json(c)
}

// 更新角色
func (*RoleController) Update(c *gin.Context) {
	var role vo.UpdateRoleVo
	if err := c.ShouldBind(&role); err != nil {
		response.NewError(err).Json(c)
		return
	}
	if ok := (&systemservice.RoleService{}).IsHasSameName(role.Name, &role.Id); ok {
		response.NewError(nil).SetMsg("角色名称已存在").Json(c)
		return
	}
	if ok := (&systemservice.RoleService{}).IsHasSameKey(role.Key, &role.Id); ok {
		response.NewError(nil).SetMsg("角色权限字符已存在").Json(c)
		return
	}
	response.New(nil, (&systemservice.RoleService{}).UpdateRole(role)).Json(c)
}

// 删除角色
func (*RoleController) Delete(c *gin.Context) {
	var ids []int
	if err := bind.BindIds(c, &ids); err != nil {
		response.NewError(err).Json(c)
		return
	}
	response.New(nil, (&systemservice.RoleService{}).DeleteRole(ids)).Json(c)
}
