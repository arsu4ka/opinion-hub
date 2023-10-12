package controllers

import (
	"net/http"

	"github.com/aru4ka/opinion-hub/internal/app/configs"

	"github.com/aru4ka/opinion-hub/internal/app/controllers/dto"
	"github.com/aru4ka/opinion-hub/internal/app/services"
	"github.com/aru4ka/opinion-hub/internal/app/utils"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type AuthController struct {
	user     *services.UserService
	jwtConf  *configs.JwtConfig
	validate *validator.Validate
}

func NewAuthController(userService *services.UserService, jwtConf *configs.JwtConfig) *AuthController {
	return &AuthController{user: userService, jwtConf: jwtConf, validate: validator.New()}
}

func (ac *AuthController) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		var createUserDto dto.CreateUserDto
		if err := c.Bind(&createUserDto); err != nil {
			return err
		}

		if err := ac.validate.Struct(createUserDto); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		if err := ac.user.CheckAvailability(createUserDto.Username, createUserDto.Email); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		if err := ac.user.Create(&createUserDto); err != nil {
			return err
		}

		return c.NoContent(http.StatusCreated)
	}
}

func (ac *AuthController) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var loginUserDto dto.LoginUserDto
		if err := c.Bind(&loginUserDto); err != nil {
			return err
		}

		if err := ac.validate.Struct(loginUserDto); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		user, err := ac.user.FindByEmail(loginUserDto.Email)
		if err != nil || !user.ComparePassword(loginUserDto.Password) {
			return c.JSON(http.StatusUnauthorized, "invalid credentials")
		}

		token, err := utils.GenerateJWT(user.ID, ac.jwtConf)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, map[string]string{"token": token})
	}
}
