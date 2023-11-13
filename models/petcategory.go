package models

import (
	"time"

	"gorm.io/gorm"
)

type PetCategory struct {
	PetCategoryId int64  `gorm:"column:pet_category_id;primaryKey;autoIncrement"`
	Name          string `gorm:"column:name"`
	Pets          []Pet
	CreatedAt     time.Time      `gorm:"autoCreateTime, column:created_at"`
	UpdatedAt     time.Time      `gorm:"autoCreateTime, column:updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at"`
}
