package routes

import (
	"github.com/aru4ka/opinion-hub/internal/app/controllers"
	"github.com/labstack/echo/v4"
)

type GlobalRouter struct {
	AuthController    *controllers.AuthController
	UserController    *controllers.UserController
	OpinionController *controllers.OpinionController
	LikeController    *controllers.LikeController
}

func (g *GlobalRouter) BindTo(base *echo.Echo) {
	authGroup := base.Group("/auth")
	authGroup.POST("/register", g.AuthController.Register())
	authGroup.POST("/login", g.AuthController.Login())

	userGroup := base.Group("/users")
	userGroup.GET("/:username", g.UserController.GetUser())
	userGroup.GET("/:username/opinions", g.OpinionController.GetUserOpinions())
	userGroup.PUT("/:username", g.UserController.UpdateUser())

	opinionGroup := base.Group("/opinions")
	opinionGroup.POST("", g.OpinionController.Create())
	opinionGroup.PUT("/:id", g.OpinionController.Update())
	opinionGroup.GET("/:id", g.OpinionController.GetOpinion())
	opinionGroup.DELETE("/:id", g.OpinionController.Delete())

	likeGroup := base.Group("/likes")
	likeGroup.POST("/:id", g.LikeController.Create())
}
