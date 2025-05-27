package systemservice

import (
	"github.com/yx1126/go-admin/DB"
	sysmodel "github.com/yx1126/go-admin/app/model/sys"
	"github.com/yx1126/go-admin/app/service"
	"github.com/yx1126/go-admin/app/vo"
)

type PostService struct{}

// 分页查询
func (*PostService) QueryPostList(params vo.PostPagingParam) (vo.PagingBackVo[vo.PostVo], error) {
	var postList []vo.PostVo
	var count int64
	query := DB.Gorm.Model(&sysmodel.SysPost{}).Order("sort,updated_at DESC,created_at DESC")
	if params.Name != "" {
		query.Where("name LIKE ?", "%"+params.Name+"%")
	}
	if params.Code != "" {
		query.Where("code LIKE ?", "%"+params.Code+"%")
	}
	if params.Status != "" {
		query.Where("status = ?", params.Status)
	}
	result := query.
		Count(&count).
		Scopes(service.PagingScope(params.Page, params.Size)).
		Find(&postList)
	return vo.PagingBackVo[vo.PostVo]{Data: postList, Count: int(count)}, result.Error
}

// 查询所有
func (*PostService) QueryPostAllList(status string) ([]vo.PostVo, error) {
	var postList []vo.PostVo
	query := DB.Gorm.Model(&sysmodel.SysPost{})
	if status != "" {
		query.Where("status = ?", status)
	}
	result := query.Find(&postList)
	return postList, result.Error
}

// 创建岗位
func (*PostService) CreatePost(post vo.CreatePostVo) error {
	return DB.Gorm.Model(&sysmodel.SysPost{}).Create(&sysmodel.SysPost{
		Name:   post.Name,
		Code:   post.Code,
		Sort:   post.Sort,
		Status: post.Status,
		Remark: post.Remark,
	}).Error
}

// 更新岗位
func (*PostService) UpdatePost(post vo.UpdatePostVo) error {
	return DB.Gorm.Model(&sysmodel.SysPost{}).
		Scopes(service.UpdateOmitScope()).
		Where("id = ?", post.Id).
		Updates(&sysmodel.SysPost{
			Name:   post.Name,
			Code:   post.Code,
			Sort:   post.Sort,
			Status: post.Status,
			Remark: post.Remark,
		}).Error
}

// 删除岗位
func (*PostService) DeletePost(ids []int) error {
	return DB.Gorm.Model(&sysmodel.SysPost{}).Delete(ids).Error
}

// 校验岗位名称
func (*PostService) HasSameName(name string, id *int) bool {
	var count int64
	query := DB.Gorm.Model(&sysmodel.SysPost{}).Where("name = ?", name)
	if id != nil {
		query.Where("id != ?", id)
	}
	query.Count(&count)
	return count > 0
}

// 校验岗位编号
func (*PostService) HasSameCode(code string, id *int) bool {
	var count int64
	query := DB.Gorm.Model(&sysmodel.SysPost{}).Where("name = ?", code)
	if id != nil {
		query.Where("id != ?", id)
	}
	query.Count(&count)
	return count > 0
}
