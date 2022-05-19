package models

import (
	"time"

	"gorm.io/gorm"
)

type Model struct {
	ID          uint           `json:"id" gorm:"primarykey"`
	CreatedAt   time.Time      `json:"createdAt"`
	CreatedByID uint           `json:"createdByID,omitempty"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	UpdatedByID uint           `json:"updatedByID,omitempty"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}
