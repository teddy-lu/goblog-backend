package models

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID int64 `gorm:"column:id;primaryKey;autoIncrement;" json:"id,omitempty"`
}

type CommonTimestampsField struct {
	CreatedAt time.Time `gorm:"column:created_at;index;" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"column:updated_at;index;" json:"updated_at,omitempty"`
}

type DeletedTimestampsField struct {
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index;" json:"deleted_at,omitempty"`
}
