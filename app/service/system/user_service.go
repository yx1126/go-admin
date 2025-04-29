package service

import (
	"github.com/yx1126/go-admin/DB"
	model "github.com/yx1126/go-admin/app/model/sys"
	"github.com/yx1126/go-admin/app/util/pwd"
	"github.com/yx1126/go-admin/app/vo"
	"github.com/yx1126/go-admin/config"
)

type UserService struct{}

func (*UserService) QueryUserList(params vo.UserQueryPageVo) ([]vo.UserVo, int, error) {
	var count int64
	var userList = make([]vo.UserVo, 0)
	query := DB.Gorm.Model(&model.SysUser{}).Select("sys_user.*", "d.dept_name").
		Omit("password").
		Joins("LEFT JOIN sys_dept as d ON sys_user.dept_id = d.id")
	if params.UserName != "" {
		query = query.Where("user_name LIKE ?", "%"+params.UserName+"%")
	}
	if params.NickName != "" {
		query = query.Where("nick_name LIKE ?", "%"+params.NickName+"%")
	}
	if params.Status != "" {
		query = query.Where("sys_user.status = ?", params.Status)
	}
	if params.DeptId != "" {
		query = query.Where("sys_user.dept_id = ?", params.DeptId)
	}
	result := query.Count(&count).
		Limit(params.Size).
		Offset((params.Page - 1) * params.Size).
		Find(&userList)
	return userList, int(count), result.Error
}

func (*UserService) QueryUserById(id int) (vo.UserVo, error) {
	var user vo.UserVo
	result := DB.Gorm.Model(&model.SysUser{}).
		Select("sys_user.*", "d.dept_name").
		Joins("LEFT JOIN sys_dept as d ON sys_user.dept_id = d.id").
		Where("sys_user.id = ?", id).
		First(&user)
	return user, result.Error
}

func (*UserService) CreateUser(user vo.CreateUserVo) error {
	pwd, _ := pwd.Encode(config.User.Password)
	return DB.Gorm.Model(&model.SysUser{}).Create(&model.SysUser{
		UserName: user.UserName,
		DeptId:   user.DeptId,
		NickName: user.NickName,
		UserType: user.UserType,
		Email:    user.Email,
		Phone:    user.Phone,
		Sex:      user.Sex,
		Avatar:   user.Avatar,
		Password: pwd,
		Status:   user.Status,
		Remark:   user.Remark,
	}).Error
}

func (*UserService) UpdateUser(user vo.UpdateUserVo) error {
	return DB.Gorm.Model(&model.SysUser{}).Where("id = ?", user.Id).Updates(&model.SysUser{
		DeptId:   user.DeptId,
		NickName: user.NickName,
		UserType: user.UserType,
		Email:    user.Email,
		Phone:    user.Phone,
		Sex:      user.Sex,
		Avatar:   user.Avatar,
		Status:   user.Status,
		Remark:   user.Remark,
	}).Error
}

func (*UserService) DeleteUser(ids []int) error {
	return DB.Gorm.Model(&model.SysUser{}).Delete(&model.SysUser{}, ids).Error
}
