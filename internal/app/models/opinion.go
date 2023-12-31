package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Opinion struct {
	ID      uuid.UUID `gorm:"primaryKey"`
	Title   string    `gorm:"type:varchar(256);not null"`
	Body    string    `gorm:"type:text"`
	IsDraft bool      `gorm:"default:true;not null"`
	Likes   uint      `gorm:"default:0"`

	CreatedAt time.Time
	UpdatedAt time.Time

	OwnerID uint
	Owner   User `gorm:"foreignKey:OwnerID"`
}

func (o *Opinion) BeforeCreate(tx *gorm.DB) error {
	if o.ID == uuid.Nil {
		o.ID = uuid.New()
	}

	return nil
}
