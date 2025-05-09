package systemservice

import (
	"github.com/yx1126/go-admin/DB"
	sysmodel "github.com/yx1126/go-admin/app/model/sys"
	"github.com/yx1126/go-admin/app/service"
	"github.com/yx1126/go-admin/app/util"
	"github.com/yx1126/go-admin/app/vo"
	"github.com/yx1126/go-admin/common/password"
	"github.com/yx1126/go-admin/config"
)

type UserService struct{}

// 查询用户列表
func (*UserService) QueryUserList(params vo.UserPagingParam) (vo.PagingBackVo[vo.UserVo], error) {
	var count int64
	var userList = make([]vo.UserVo, 0)
	query := DB.Gorm.Model(&sysmodel.SysUser{}).
		Select("sys_user.*", "d.name as dept_name").
		Omit("password").
		Joins("LEFT JOIN sys_dept as d ON sys_user.dept_id = d.id")
	if params.UserName != "" {
		query = query.Where("sys_user.user_name LIKE ?", "%"+params.UserName+"%")
	}
	if params.NickName != "" {
		query = query.Where("sys_user.nick_name LIKE ?", "%"+params.NickName+"%")
	}
	if params.Status != "" {
		query = query.Where("sys_user.status = ?", params.Status)
	}
	if params.DeptId != "" {
		query = query.Where("sys_user.dept_id = ?", params.DeptId)
	}
	result := query.
		Count(&count).
		Scopes(service.PagingScope(params.Page, params.Size)).
		Find(&userList)
	return vo.PagingBackVo[vo.UserVo]{Data: userList, Count: int(count)}, result.Error
}

// 查询所有用户列表
func (*UserService) QueryUserAllList(params vo.UserParam) ([]vo.UserVo, error) {
	var userList = make([]vo.UserVo, 0)
	query := DB.Gorm.Model(&sysmodel.SysUser{}).Select("sys_user.*", "d.name as deptName").
		Omit("password").
		Joins("LEFT JOIN sys_dept as d ON sys_user.dept_id = d.id")
	if params.UserName != "" {
		query = query.Where("sys_user.user_name LIKE ?", "%"+params.UserName+"%")
	}
	if params.NickName != "" {
		query = query.Where("sys_user.nick_name LIKE ?", "%"+params.NickName+"%")
	}
	if params.Status != "" {
		query = query.Where("sys_user.status = ?", params.Status)
	}
	if params.DeptId != "" {
		query = query.Where("sys_user.dept_id = ?", params.DeptId)
	}
	result := query.Find(&userList)
	return userList, result.Error
}

// 根据id查询用户信息
func (*UserService) QueryUserById(id int) (*vo.UserInfoVo, error) {
	var user vo.UserInfoVo
	query := DB.Gorm.Model(&sysmodel.SysUser{}).
		Select("sys_user.*", "d.name").
		Joins("LEFT JOIN sys_dept as d ON sys_user.dept_id = d.id").
		Where("sys_user.id = ?", id)
	if err := query.First(&user.UserVo).Error; err != nil {
		return nil, err
	}

	if err := DB.Gorm.Model(&sysmodel.SysUserPost{}).
		Select("post_id").
		Where("user_id = ?", id).
		Pluck("post_id", &user.PostIds).Error; err != nil {
		return nil, err
	}
	if err := DB.Gorm.Model(&sysmodel.SysUserRole{}).
		Select("role_id").
		Where("user_id = ?", id).
		Pluck("role_id", &user.RoleIds).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// 创建用户
func (*UserService) CreateUser(user vo.CreateUserVo) error {
	tx := DB.Gorm.Begin()
	pwd, _ := password.Encode(config.User.Password)
	// 插入用户
	sysUser := sysmodel.SysUser{
		UserName: user.UserName,
		DeptId:   user.DeptId,
		NickName: user.NickName,
		UserType: user.UserType,
		Email:    user.Email,
		Phone:    user.Phone,
		Sex:      user.Sex,
		Password: pwd,
		Status:   user.Status,
		Remark:   user.Remark,
	}
	result := tx.Model(&sysmodel.SysUser{}).Create(&sysUser)

	if err := result.Error; err != nil {
		tx.Rollback()
		return err
	}
	// 插入岗位
	if user.PostIds != nil && len(*user.PostIds) > 0 {
		postList := util.Map(*user.PostIds, func(item, _ int) sysmodel.SysUserPost {
			return sysmodel.SysUserPost{
				UserId: sysUser.Id,
				PostId: item,
			}
		})
		if err := tx.Model(&sysmodel.SysUserPost{}).Create(&postList).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	// 插入角色
	if user.RoleIds != nil && len(*user.RoleIds) > 0 {
		roleList := util.Map(*user.RoleIds, func(item, _ int) sysmodel.SysUserRole {
			return sysmodel.SysUserRole{
				UserId: sysUser.Id,
				RoleId: item,
			}
		})
		if err := tx.Model(&sysmodel.SysUserRole{}).Create(&roleList).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit().Error
}

// 更新用户
func (*UserService) UpdateUser(user vo.UpdateUserVo) error {
	tx := DB.Gorm.Begin()
	if err := tx.Model(&sysmodel.SysUser{}).
		Select("dept_id", "nick_name", "user_type", "email", "phone", "sex", "avatar", "status", "remark").
		Where("id = ?", user.Id).
		Updates(&sysmodel.SysUser{
			DeptId:   user.DeptId,
			NickName: user.NickName,
			UserType: user.UserType,
			Email:    user.Email,
			Phone:    user.Phone,
			Sex:      user.Sex,
			Status:   user.Status,
			Remark:   user.Remark,
		}).Error; err != nil {
		tx.Rollback()
		return err
	}
	// 删除岗位
	if err := tx.Model(&sysmodel.SysUserPost{}).Where("user_id = ?", user.Id).Delete(&sysmodel.SysUserPost{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	// 删除角色
	if err := tx.Model(&sysmodel.SysUserRole{}).Where("user_id = ?", user.Id).Delete(&sysmodel.SysUserRole{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	// 插入岗位
	if user.PostIds != nil && len(*user.PostIds) > 0 {
		postList := util.Map(*user.PostIds, func(item, _ int) sysmodel.SysUserPost {
			return sysmodel.SysUserPost{
				UserId: user.Id,
				PostId: item,
			}
		})
		if err := tx.Model(&sysmodel.SysUserPost{}).Create(&postList).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	// 插入角色
	if user.RoleIds != nil && len(*user.RoleIds) > 0 {
		roleList := util.Map(*user.RoleIds, func(item, _ int) sysmodel.SysUserRole {
			return sysmodel.SysUserRole{
				UserId: user.Id,
				RoleId: item,
			}
		})
		if err := tx.Model(&sysmodel.SysUserRole{}).Create(&roleList).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit().Error
}

// 删除用户
func (*UserService) DeleteUser(ids []int) error {
	return DB.Gorm.Model(&sysmodel.SysUser{}).Delete(&sysmodel.SysUser{}, ids).Error
}

// 修改密码
func (*UserService) UpdatePwd(id int, pwd string) error {
	return DB.Gorm.Model(&sysmodel.SysUser{}).Where("id = ?", id).Update("password", pwd).Error
}

// 校验用户名
func (*UserService) IsHasSameName(name string) bool {
	var count int64
	DB.Gorm.Model(&sysmodel.SysUser{}).Where("user_name = ?", name).Count(&count)
	return count > 0
}
