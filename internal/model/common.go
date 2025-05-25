package model

import (
	"gorm.io/gorm"
	"time"
)

type CommonModel struct {
	ID        int            `gorm:"primary_key" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
