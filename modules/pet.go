package modules

import (
	"time"

	"gorm.io/gorm"
)

type Pet struct {
	PetId         int64          `gorm: "column:pet_id;primaryKey"`
	Name          string         `gorm: "column:name"`
	DateOfBirth   time.Time      `gorm: "column:date_of_birth"`
	IsAggresive   bool           `gorm: "column:is_aggresive"`
	PetCategoryId int64          `gorm: "column:pet_category_id"`
	Owners        []Owner        `gorm: "foreignKey:PetId"`
	CreatedAt     time.Time      `gorm: "autoCreateTime, column:created_at"`
	UpdatedAt     time.Time      `gorm: "autoCreateTime, column:updated_at"`
	DeletedAt     gorm.DeletedAt `gorm: "column:deleted_at"`
}
