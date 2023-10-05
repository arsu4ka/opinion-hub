package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"golang.org/x/crypto/bcrypt"
)

func passwordValidator(password string) bool {
	if len(password) < 6 || len(password) > 30 {
		return false
	}
	return true
}

func nameValidator(name string) bool {
	rule1 := govalidator.IsAlpha(name)
	rule2 := name != ""
	return rule1 && rule2
}

type User struct {
	ID        uint
	FirstName string `gorm:"type:varchar(255);not null"`
	LastName  string `gorm:"type:varchar(255)"`
	Username  string `gorm:"type:varchar(255);unique;not null"`
	Email     string `gorm:"type:varchar(255);unique;not null"`
	IsActive  bool   `gorm:"default:false;not null"`
	IsPublic  bool   `gorm:"not null"`
	Password  string `gorm:"type:varchar(30);not null"`
	CreatedAt time.Time
}

func (u *User) Validate() error {
	return validation.ValidateStruct(
		u,
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Username, validation.Required, is.Alphanumeric),
		validation.Field(&u.Password, validation.NewStringRule(
			passwordValidator,
			"invalid password",
		)),
		validation.Field(&u.FirstName, validation.NewStringRule(
			nameValidator,
			"invalid first name",
		)),
		validation.Field(&u.LastName, validation.NewStringRule(
			nameValidator,
			"invalid last name",
		)),
		validation.Field(&u.IsPublic, validation.Required),
	)
}

func (u *User) HashPassword() error {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), 12)
	if err != nil {
		return err
	}

	u.Password = string(hashedBytes)
	return nil
}

func (u *User) ComparePassword(compareTo string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(compareTo))
	return err == nil
}
