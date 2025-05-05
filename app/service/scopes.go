package service

import (
	"strings"

	"github.com/yx1126/go-admin/app/util"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

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

// 默认排序
func SortScope(DB *gorm.DB) *gorm.DB {
	return DB.Order(clause.OrderBy{
		Columns: []clause.OrderByColumn{
			{Column: clause.Column{Name: "sort"}, Desc: false},
			{Column: clause.Column{Name: "updated_at"}, Desc: true},
			{Column: clause.Column{Name: "created_at"}, Desc: true},
		},
	})
}

// 默认排序
func SortBeforeScope(columns ...string) func(DB *gorm.DB) *gorm.DB {
	base_columns := []clause.OrderByColumn{}
	for _, v := range columns {
		desc := false
		strs := strings.Split(v, ",")
		if descStr, _ := util.SliceAt(strs, 1); descStr == "desc" {
			desc = true
		}
		if name, _ := util.SliceAt(strs, 0); name != "" {
			base_columns = append(base_columns, clause.OrderByColumn{
				Column: clause.Column{Name: name},
				Desc:   desc,
			})
		}
	}
	base_columns = append(base_columns, []clause.OrderByColumn{
		{Column: clause.Column{Name: "sort"}, Desc: false},
		{Column: clause.Column{Name: "updated_at"}, Desc: true},
		{Column: clause.Column{Name: "created_at"}, Desc: true},
	}...)
	return func(DB *gorm.DB) *gorm.DB {
		return DB.Order(clause.OrderBy{
			Columns: base_columns,
		})
	}
}

// 默认排序
func SortAfterScope(columns ...string) func(DB *gorm.DB) *gorm.DB {
	base_columns := []clause.OrderByColumn{
		{Column: clause.Column{Name: "sort"}, Desc: false},
		{Column: clause.Column{Name: "updated_at"}, Desc: true},
		{Column: clause.Column{Name: "created_at"}, Desc: true},
	}
	for _, v := range columns {
		desc := false
		strs := strings.Split(v, ",")
		if descStr, _ := util.SliceAt(strs, 1); descStr == "desc" {
			desc = true
		}
		if name, _ := util.SliceAt(strs, 0); name != "" {
			base_columns = append(base_columns, clause.OrderByColumn{
				Column: clause.Column{Name: name},
				Desc:   desc,
			})
		}
	}
	return func(DB *gorm.DB) *gorm.DB {
		return DB.Order(clause.OrderBy{
			Columns: base_columns,
		})
	}
}
