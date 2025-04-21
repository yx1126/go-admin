package model

import (
	"github.com/yx1126/go-admin/app/util/datetime"
	"gorm.io/gorm"
)

type BaseModel struct {
	CreatedAt datetime.Datetime `json:"createdAt"`
	UpdatedAt datetime.Datetime `json:"updatedAt"`
	CreatedBy string            `json:"createdBy"`
	UpdatedBy string            `json:"updatedBy"`
	DeletedAt gorm.DeletedAt    `json:"-" gorm:"index"`
}
