package controllers

import (
	"errors"
	"net/http"

	"github.com/aru4ka/opinion-hub/internal/app/controllers/dto"
	"github.com/aru4ka/opinion-hub/internal/app/services"
	"github.com/aru4ka/opinion-hub/internal/app/utils"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type UserController struct {
	user     *services.UserService
	opinion  *services.OpinionService
	validate *validator.Validate
}

func NewUserController(userService *services.UserService, opinionService *services.OpinionService) *UserController {
	return &UserController{
		user:     userService,
		opinion:  opinionService,
		validate: validator.New(),
	}
}

func (uc *UserController) GetUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		username := c.Param("username")

		user, err := uc.user.FindByUsername(username)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return echo.ErrNotFound
			}
			return err
		}

		return c.JSON(http.StatusOK, dto.NewResponseUserDto(user).HideEmail())
	}
}

func (uc *UserController) UpdateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		username := c.Param("username")
		token := c.Get("user").(*jwt.Token)
		claims := token.Claims.(*utils.JwtCustomClaims)

		user, err := uc.user.FindByUsername(username)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return echo.ErrNotFound
			}
			return err
		}

		if user.ID != claims.UserId {
			return echo.ErrForbidden
		}

		var updateUserDto dto.UpdateUserDto
		if err := c.Bind(&updateUserDto); err != nil {
			return err
		}

		if err := uc.validate.Struct(updateUserDto); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		if !user.ComparePassword(updateUserDto.Password) {
			return c.JSON(http.StatusBadRequest, "invalid password")
		}

		if err := uc.user.CheckAvailability(updateUserDto.Username, updateUserDto.Email); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		if err := uc.user.Update(user.ID, &updateUserDto); err != nil {
			return err
		}

		return c.NoContent(http.StatusOK)
	}
}

func (uc *UserController) GetOpinions() echo.HandlerFunc {
	return func(c echo.Context) error {
		username := c.Param("username")

		user, err := uc.user.FindByUsername(username)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return echo.ErrNotFound
			}
			return err
		}

		if !user.IsPublic {
			return c.JSON(http.StatusForbidden, []*dto.ResponseOpinionDto{})
		}

		opinions, err := uc.opinion.FindByUserID(user.ID, false)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		opinionsDto := make([]*dto.ResponseOpinionDto, len(opinions))
		for i, op := range opinions {
			opinionsDto[i] = dto.NewResponseOpinionDto(op)
		}

		return c.JSON(http.StatusOK, opinionsDto)
	}
}
