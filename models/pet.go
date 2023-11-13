package models

import (
	"time"

	"gorm.io/gorm"
)

type Pet struct {
	PetId         int64          `gorm:"column:pet_id;primaryKey;autoIncrement"`
	Name          string         `gorm:"column:name"`
	DateOfBirth   time.Time      `gorm:"column:date_of_birth"`
	IsAggresive   bool           `gorm:"column:is_aggressive"`
	PetCategoryId int64          `gorm:"column:pet_category_id"`
	PetCategory   PetCategory    `gorm:"foreignKey:pet_category_id;references:pet_category_id;not null"`
	CreatedAt     time.Time      `gorm:"autoCreateTime,column:created_at"`
	UpdatedAt     time.Time      `gorm:"autoCreateTime,column:updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at"`
}
