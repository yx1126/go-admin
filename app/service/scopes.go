package service

import "gorm.io/gorm"

// 更新0值时 排除默认字段
func UpdateOmitScope(columns ...string) func(db *gorm.DB) *gorm.DB {
	defOmit := []string{"id", "created_at", "update_at", "deleted_at"}
	defOmit = append(defOmit, columns...)
	return func(db *gorm.DB) *gorm.DB {
		return db.Select("*").Omit(defOmit...)
	}
}
