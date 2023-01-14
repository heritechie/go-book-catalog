package models

import (
	"time"

	"github.com/google/uuid"
)

type Book struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Title       string    `gorm:"type:varchar(255);not null"`
	Description string    `gorm:"type:varchar(255);not null"`
	Year        string    `gorm:"type:varchar(255);not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	AuthorID    uuid.UUID  `gorm:"type:uuid REFERENCES author(id)"`
	PublisherID uuid.UUID  `gorm:"type:uuid REFERENCES publisher(id)"`
	Author      *Author    `gorm:"foreignKey:AuthorID;associationForeignKey:ID,constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Publisher   *Publisher `gorm:"foreignKey:PublisherID;associationForeignKey:ID,constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
