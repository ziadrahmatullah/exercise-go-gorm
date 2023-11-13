package modules

import (
	"time"

	"gorm.io/gorm"
)

type PetCategory struct {
	PetCategoryId int64          `gorm: "column:pet_category_id;primaryKey"`
	Name          string         `gorm: "column:name"`
	Pets          []Pet          `gorm: "foreignKey:PetCategoryId`
	CreatedAt     time.Time      `gorm: "autoCreateTime, column:created_at"`
	UpdatedAt     time.Time      `gorm: "autoCreateTime, column:updated_at"`
	DeletedAt     gorm.DeletedAt `gorm: "column:deleted_at"`
}
