package controllers

import (
	"errors"
	"net/http"

	"github.com/aru4ka/opinion-hub/internal/app/services"
	"github.com/aru4ka/opinion-hub/internal/app/utils"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type LikeController struct {
	opinion *services.OpinionService
	like    *services.LikeService
}

func NewLikeController(likeService *services.LikeService, opinionService *services.OpinionService) *LikeController {
	return &LikeController{like: likeService, opinion: opinionService}
}

func (lc *LikeController) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user").(*jwt.Token)
		claims := token.Claims.(*utils.JwtCustomClaims)

		opinionId, err := uuid.Parse(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
		}

		if _, err := lc.opinion.FindById(opinionId); err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return echo.ErrNotFound
			}
			return err
		}

		if err := lc.like.Create(claims.UserId, opinionId); err != nil {
			return err
		}

		return c.NoContent(http.StatusCreated)
	}
}
