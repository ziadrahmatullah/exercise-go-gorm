package models

import (
	"time"

	"gorm.io/gorm"
)

type Owner struct {
	OwnerId   int64          `gorm:"column:owner_id;primaryKey;autoIncrement"`
	Name      string         `gorm:"column:name"`
	Address   string         `gorm:"column:address"`
	Pets      []Pet          `gorm:"many2many:owner_pets"`
	CreatedAt time.Time      `gorm:"autoCreateTime,column:created_at"`
	UpdatedAt time.Time      `gorm:"autoCreateTime,column:updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
}
