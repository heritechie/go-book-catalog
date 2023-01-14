package models

import (
	"time"

	"github.com/google/uuid"
)

type Publisher struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;"`
	Name      string    `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
