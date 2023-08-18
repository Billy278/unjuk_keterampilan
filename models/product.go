package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	Id        uint64  `json:"id"`
	Name      string  `json:"name" validate:"required"`
	Prize     float64 `json:"prize" validate:"required"`
	Stock     uint64  `json:"stock" validate:"required,min=10"`
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeleteAt  gorm.DeletedAt
}
