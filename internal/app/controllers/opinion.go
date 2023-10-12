package controllers

import (
	"errors"
	"net/http"

	"github.com/aru4ka/opinion-hub/internal/app/controllers/dto"
	"github.com/aru4ka/opinion-hub/internal/app/services"
	"github.com/aru4ka/opinion-hub/internal/app/utils"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type OpinionController struct {
	opinion  *services.OpinionService
	validate *validator.Validate
}

func NewOpinionController(opinionService *services.OpinionService) *OpinionController {
	return &OpinionController{opinion: opinionService, validate: validator.New()}
}

func (oc *OpinionController) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user").(*jwt.Token)
		claims := token.Claims.(*utils.JwtCustomClaims)

		opDto := new(dto.CreateOpinionDto)
		if err := c.Bind(opDto); err != nil {
			return err
		}

		if err := oc.validate.Struct(opDto); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		if err := oc.opinion.Create(claims.UserId, opDto); err != nil {
			return err
		}

		return c.NoContent(http.StatusCreated)
	}
}

func (oc *OpinionController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user").(*jwt.Token)
		claims := token.Claims.(*utils.JwtCustomClaims)

		opDto := new(dto.UpdateOpinionDto)
		if err := c.Bind(opDto); err != nil {
			return err
		}

		if err := oc.validate.Struct(opDto); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		id, err := uuid.Parse(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		opinion, err := oc.opinion.FindByID(id)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return c.NoContent(http.StatusNotFound)
			}
			return err
		}

		if opinion.OwnerID != claims.UserId {
			return c.NoContent(http.StatusForbidden)
		}

		if err := oc.opinion.Update(id, claims.UserId, opDto); err != nil {
			return err
		}

		return c.NoContent(http.StatusOK)
	}
}
