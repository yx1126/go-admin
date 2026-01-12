package model

import (
	"go-admin/common/types"

	"gorm.io/gorm"
)

type BaseModel struct {
	Id        int            `json:"id,string" gorm:"primaryKey;autoIncrement"`
	CreatedAt types.Datetime `json:"createdAt"`
	UpdatedAt types.Datetime `json:"updatedAt"`
	CreatedBy string         `json:"createdBy"`
	UpdatedBy string         `json:"updatedBy"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
