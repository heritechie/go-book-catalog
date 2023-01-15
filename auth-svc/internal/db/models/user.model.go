package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name      *string   `gorm:"type:varchar(255)"`
	Dob       *string   `gorm:"type:varchar(255)"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
