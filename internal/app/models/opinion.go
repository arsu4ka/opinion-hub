package models

import (
	"errors"
	"github.com/google/uuid"
	"time"
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

func (o *Opinion) Validate() error {
	if o.ID == uuid.Nil {
		return errors.New("opinion id is required")
	}
	if o.Title == "" {
		return errors.New("opinion title is required")
	}

	return nil
}
