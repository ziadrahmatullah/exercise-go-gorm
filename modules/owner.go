package modules

import (
	"time"

	"gorm.io/gorm"
)

type Owner struct {
	OwnerId   int64          `gorm: "column:owner_id;primaryKey"`
	Name      string         `gorm: "column:name"`
	Address   string         `gorm: "column:address"`
	OwnerPets []OwnerPet     `gorm: "foreignKey:OwnerId"`
	CreatedAt time.Time      `gorm: "autoCreateTime, column:created_at"`
	UpdatedAt time.Time      `gorm: "autoCreateTime, column:updated_at"`
	DeletedAt gorm.DeletedAt `gorm: "column:deleted_at"`
}
