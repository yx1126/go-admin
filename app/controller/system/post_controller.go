package systemcontroller

import (
	"github.com/gin-gonic/gin"
	systemservice "github.com/yx1126/go-admin/app/service/system"
	"github.com/yx1126/go-admin/app/vo"
	"github.com/yx1126/go-admin/common/constant"
	bind "github.com/yx1126/go-admin/common/should_bind"
	"github.com/yx1126/go-admin/response"
)

type PostController struct{}

// 分页查询
func (*PostController) QueryPostList(c *gin.Context) {
	var params vo.PostPagingParam
	if err := bind.BindPaging(c, &params); err != nil {
		response.NewError(err).Json(c)
		return
	}
	data, err := (&systemservice.PostService{}).QueryPostList(params)
	paging := response.Paging{
		Page:  params.Page,
		Size:  params.Size,
		Total: data.Count,
		List:  data.Data,
	}
	response.New(nil, err).SetPaging(paging).Json(c)
}

// 查询未禁用的所有岗位
func (*PostController) QueryPostAllList(c *gin.Context) {
	response.New((&systemservice.PostService{}).QueryPostAllList(constant.STATUS)).Json(c)
}

// 新增岗位
func (*PostController) Create(c *gin.Context) {
	var post vo.CreatePostVo
	if err := bind.ShouldBindJSON(c, &post); err != nil {
		response.NewError(err).Json(c)
		return
	}
	if ok := (&systemservice.PostService{}).HasSameName(post.Name, nil); ok {
		response.NewError(nil).SetMsg("岗位名称已存在").Json(c)
		return
	}
	if ok := (&systemservice.PostService{}).HasSameCode(post.Code, nil); ok {
		response.NewError(nil).SetMsg("岗位编号已存在").Json(c)
		return
	}
	response.New(nil, (&systemservice.PostService{}).CreatePost(post)).Json(c)
}

// 更新岗位
func (*PostController) Update(c *gin.Context) {
	var post vo.UpdatePostVo
	if err := bind.ShouldBindJSON(c, &post); err != nil {
		response.NewError(err).Json(c)
		return
	}
	if ok := (&systemservice.PostService{}).HasSameName(post.Name, &post.Id); ok {
		response.NewError(nil).SetMsg("岗位名称已存在").Json(c)
		return
	}
	if ok := (&systemservice.PostService{}).HasSameCode(post.Code, &post.Id); ok {
		response.NewError(nil).SetMsg("岗位编号已存在").Json(c)
		return
	}
	response.New(nil, (&systemservice.PostService{}).UpdatePost(post)).Json(c)
}

// 删除用户
func (*PostController) Delete(c *gin.Context) {
	var ids []int
	if err := bind.BindIds(c, &ids); err != nil {
		response.NewError(err).Json(c)
		return
	}
	response.New(nil, (&systemservice.PostService{}).DeletePost(ids)).Json(c)
}
