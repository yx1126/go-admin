package service

import "gorm.io/gorm"

// 更新0值时 排除默认字段
func UpdateOmitScope(columns ...string) func(DB *gorm.DB) *gorm.DB {
	defOmit := []string{"id", "created_at", "update_at", "deleted_at"}
	defOmit = append(defOmit, columns...)
	return func(DB *gorm.DB) *gorm.DB {
		return DB.Select("*").Omit(defOmit...)
	}
}

// 分页
func PagingScope(page, size int) func(DB *gorm.DB) *gorm.DB {
	return func(DB *gorm.DB) *gorm.DB {
		return DB.Offset((page - 1) * size).Limit(size)
	}
}
