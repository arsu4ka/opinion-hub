package dto

import (
	"github.com/aru4ka/opinion-hub/internal/app/models"
)

type CreateUserDto struct {
	FirstName string `json:"firstName" xml:"firstName" validate:"required,alpha,min=2"`
	LastName  string `json:"lastName" xml:"lastName" validate:"omitempty,alpha,min=2"`
	Username  string `json:"username" xml:"username" validate:"required,alphanum,min=4"`
	Email     string `json:"email" xml:"email" validate:"required,email"`
	IsPublic  bool   `json:"isPublic" xml:"isPublic" validate:"required"`
	Password  string `json:"password" xml:"password" validate:"required,min=6,max=30"`
}

func (ud *CreateUserDto) ToModel() *models.User {
	return &models.User{
		FirstName: ud.FirstName,
		LastName:  ud.LastName,
		Username:  ud.Username,
		Email:     ud.Email,
		IsPublic:  ud.IsPublic,
		Password:  ud.Password,
	}
}

type UpdateUserDto struct {
	CreateUserDto
}

type LoginUserDto struct {
	Email    string `json:"email" xml:"email" validate:"required,email"`
	Password string `json:"password" xml:"password" validate:"required,min=6,max=30"`
}

type ResponseUserDto struct {
	FirstName string `json:"firstName" xml:"firstName"`
	LastName  string `json:"lastName" xml:"lastName"`
	Username  string `json:"username" xml:"username"`
	Email     string `json:"email" xml:"email"`
	IsPublic  bool   `json:"isPublic" xml:"isPublic"`
}

func NewResponseUserDto(user *models.User) *ResponseUserDto {
	return &ResponseUserDto{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
		Email:     user.Email,
		IsPublic:  user.IsPublic,
	}
}

func (r *ResponseUserDto) HideEmail() *ResponseUserDto {
	return &ResponseUserDto{
		FirstName: r.FirstName,
		LastName:  r.LastName,
		Username:  r.Username,
		IsPublic:  r.IsPublic,
	}
}
