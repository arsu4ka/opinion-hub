package models

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Like struct {
	ID uuid.UUID `gorm:"primaryKey"`

	UserID uint
	User   User `gorm:"foreignKey:UserID"`

	OpinionID uuid.UUID
	Opinion   Opinion `gorm:"foreignKey:OpinionID"`
}

func (l *Like) Validate() error {
	if l.ID == uuid.Nil {
		return errors.New("like id is required")
	}
	return nil
}

func (l *Like) BeforeCreate(tx *gorm.DB) (err error) {
	if l.ID == uuid.Nil {
		l.ID = uuid.New()
	}

	opinion := &Opinion{ID: l.OpinionID}
	if err := tx.First(opinion).Error; err != nil {
		return err
	}

	opinion.Likes++
	return tx.Save(opinion).Error
}
