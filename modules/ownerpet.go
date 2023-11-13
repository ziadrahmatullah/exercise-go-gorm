package modules

import (
	"time"

	"gorm.io/gorm"
)

type OwnerPet struct {
	OwnerPetId int64          `gorm: "column:owner_pet_id;primaryKey"`
	OwnerId    int64          `gorm: "column:owner_id"`
	PetId      int64          `gorm: "column:pet_id"`
	StartDate  time.Time      `gorm: "column:start_date"`
	EndDate    time.Time      `gorm: "column:end_date" `
	CreatedAt  time.Time      `gorm: "autoCreateTime, column:created_at"`
	UpdatedAt  time.Time      `gorm: "autoCreateTime, column:updated_at"`
	DeletedAt  gorm.DeletedAt `gorm: "column:deleted_at"`
}
