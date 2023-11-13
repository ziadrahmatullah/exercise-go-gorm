package models

import (
	"time"

	"gorm.io/gorm"
)

type OwnerPet struct {
	OwnerPetId int64          `gorm:"column:owner_pet_id;primaryKey;autoIncrement"`
	OwnerId    int64          `gorm:"column:owner_id"`
	Owner      Owner          `gorm:"foreignKey:owner_id;references:owner_id;not null"`
	PetId      int64          `gorm:"column:pet_id"`
	Pet        Pet            `gorm:"foreignKey:pet_id;references:pet_id;not null"`
	StartDate  *time.Time     `gorm:"column:start_date"`
	EndDate    *time.Time     `gorm:"column:end_date"`
	CreatedAt  time.Time      `gorm:"autoCreateTime,column:created_at"`
	UpdatedAt  time.Time      `gorm:"autoCreateTime,column:updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at"`
}
