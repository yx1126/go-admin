package model

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	CreateBy   string
	CreateTime time.Time
	UpdateBy   string
	UpdateTime time.Time
	DeleteTime gorm.DeletedAt
}
