package models

import (
	"time"

	"github.com/google/uuid"
)

type UserAuth struct {
	Username       string `gorm:"type:varchar(255);primary_key;"`
	HashedPassword string `gorm:"type:varchar(255);not null"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	UserID         uuid.UUID
	User           User `gorm:"references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
