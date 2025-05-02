package systemservice

import (
	"github.com/yx1126/go-admin/DB"
	sysmodel "github.com/yx1126/go-admin/app/model/sys"
	"github.com/yx1126/go-admin/app/service"
	"github.com/yx1126/go-admin/app/util"
	"github.com/yx1126/go-admin/app/vo"
)

type DeptService struct{}

// 查询部门列表
func (*DeptService) QueryDeptList(param vo.DeptParam) ([]vo.DeptTreeVo, error) {
	deptList := make([]vo.DeptTreeVo, 0)
	query := DB.Gorm.Model(&sysmodel.SysDept{}).
		Select("sys_dept.*", "u.user_name as leader_name", "u.nick_name as leader_nick_name", "u.email", "u.phone").
		Joins("LEFT JOIN sys_user as u ON u.id = sys_dept.leader_id").
		Order("parent_id,sort,id")
	if param.Name != "" {
		query.Where("name LIKE ?", "%"+param.Name+"%")
	}
	if param.Status != "" {
		query.Where("status = ?", param.Status)
	}
	result := query.Find(&deptList)
	return deptList, result.Error
}

// 查询部门树结构
func (*DeptService) QueryDeptSelectTree() ([]vo.DeptVo, error) {
	deptList := make([]vo.DeptVo, 0)
	result := DB.Gorm.Model(&sysmodel.SysDept{}).Order("parent_id,sort,id").Find(&deptList)
	return util.ListToTree(deptList, 0), result.Error
}

// 创建部门
func (*DeptService) CreateDept(dept vo.CreateDeptVo) error {
	return DB.Gorm.Model(&sysmodel.SysDept{}).Create(&sysmodel.SysDept{
		ParentId: dept.ParentId,
		Name:     dept.Name,
		Sort:     dept.Sort,
		LeaderId: dept.LeaderId,
		Status:   dept.Status,
	}).Error
}

// 修改部门
func (*DeptService) UpdateDept(dept vo.UpdateDeptVo) error {
	return DB.Gorm.Model(&sysmodel.SysDept{}).
		Scopes(service.UpdateOmitScope()).
		Where("id = ?", dept.Id).Updates(&sysmodel.SysDept{
		ParentId: dept.ParentId,
		Name:     dept.Name,
		Sort:     dept.Sort,
		LeaderId: dept.LeaderId,
		Status:   dept.Status,
	}).Error
}

// 删除部门
func (*DeptService) DeleteDept(ids []int) error {
	return DB.Gorm.Delete(&sysmodel.SysDept{}, ids).Error
}

// 根据id查询时候是否有子集
func (*DeptService) DeptHasChildren(parentId int) bool {
	var count int64
	result := DB.Gorm.Model(&sysmodel.SysDept{}).Where("parent_id = ?", parentId).Count(&count)
	if result.Error != nil {
		count = 0
	}
	return count > 0
}

// 校验部门名称
func (*DeptService) DeptHasSameName(name string, id *int) bool {
	var count int64
	query := DB.Gorm.Model(&sysmodel.SysDept{}).Where("name = ?", name)
	if id != nil {
		query.Where("id != ?", id)
	}
	if result := query.Count(&count); result.Error != nil {
		count = 0
	}
	return count > 0
}
