package models

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	ID          uint `gorm:"primaryKey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Description string         `gorm:"size:150;not_null" json:"description,omitempty" binding:"required"`
	IsCompleted bool           `gorm:"default:false;not_null" json:"is_completed"`
}
