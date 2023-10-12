package routes

import (
	"github.com/aru4ka/opinion-hub/internal/app/controllers"
	"github.com/labstack/echo/v4"
)

type GlobalRouter struct {
	AuthController *controllers.AuthController
	UserController *controllers.UserController
}

func (g *GlobalRouter) BindTo(base *echo.Echo) {
	authGroup := base.Group("/auth")
	authGroup.POST("/register", g.AuthController.Register())
	authGroup.POST("/login", g.AuthController.Login())

	userGroup := base.Group("/users")
	userGroup.GET("/:username", g.UserController.GetUser())
	userGroup.GET("/:username/opinions", g.UserController.GetOpinions())
	userGroup.PUT("/:username", g.UserController.UpdateUser())
}