package model

import (
	"github.com/yx1126/go-admin/app/util/datetime"
	"gorm.io/gorm"
)

type BaseModel struct {
	gorm.Model
	CreatedAt datetime.Datetime
	UpdatedAt datetime.Datetime
	CreatedBy string `json:"createBy"`
	UpdatedBy string `json:"updateBy"`
}
