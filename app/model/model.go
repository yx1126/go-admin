package model

import (
	"gorm.io/gorm"
)

type BaseModel struct {
	gorm.Model
	CreatedBy string `json:"createBy"`
	UpdatedBy string `json:"updateBy"`
}
