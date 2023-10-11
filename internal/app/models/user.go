package models

import (
	"time"

	"gorm.io/gorm"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"golang.org/x/crypto/bcrypt"
)

func getPasswordValidator(isHashed bool) func(string) bool {
	return func(password string) bool {
		if isHashed {
			return true
		}
		return len(password) >= 6 && len(password) <= 30
	}
}

type User struct {
	ID        uint
	FirstName string `gorm:"type:varchar(255);not null"`
	LastName  string `gorm:"type:varchar(255)"`
	Username  string `gorm:"type:varchar(255);unique;not null"`
	Email     string `gorm:"type:varchar(255);unique;not null"`
	IsPublic  bool   `gorm:"not null"`
	Password  string `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time

	isHashedPassword bool `gorm:"-"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	return u.HashPassword()
}

func (u *User) SetHashed() {
	u.isHashedPassword = true
}

func (u *User) Validate() error {
	return validation.ValidateStruct(
		u,
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Username, validation.Required, is.Alphanumeric),
		validation.Field(&u.Password, validation.Required, validation.NewStringRule(
			getPasswordValidator(u.isHashedPassword),
			"invalid password",
		)),
		validation.Field(&u.FirstName, validation.Required),
		validation.Field(&u.LastName),
		validation.Field(&u.IsPublic, validation.Required),
	)
}

func (u *User) HashPassword() error {
	if u.isHashedPassword {
		return nil
	}

	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), 12)
	if err != nil {
		return err
	}

	u.Password = string(hashedBytes)
	u.isHashedPassword = true
	return nil
}

func (u *User) ComparePassword(compareTo string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(compareTo))
	return err == nil
}
