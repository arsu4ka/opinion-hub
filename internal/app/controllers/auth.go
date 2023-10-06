package controllers

import (
	"errors"
	"github.com/aru4ka/opinion-hub/internal/app/controllers/dto"
	"github.com/aru4ka/opinion-hub/internal/app/services"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

type AuthController struct {
	user     *services.UserService
	validate *validator.Validate
}

func NewAuthController(userService *services.UserService) *AuthController {
	return &AuthController{user: userService, validate: validator.New()}
}

func (ac *AuthController) SignUp() echo.HandlerFunc {
	return func(c echo.Context) error {
		var createUserDto dto.CreateUserDto
		if err := c.Bind(&createUserDto); err != nil {
			return err
		}

		if err := ac.validate.Struct(createUserDto); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		user := createUserDto.ToModel()
		if err := user.HashPassword(); err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		if err := ac.user.Create(user); err != nil {
			if errors.Is(err, gorm.ErrDuplicatedKey) {
				return c.JSON(http.StatusBadRequest, err.Error())
			}
			return err
		}

		return c.JSON(http.StatusCreated, dto.NewResponseUserDto(user))
	}
}
