package server

import (
	"github.com/aru4ka/opinion-hub/internal/app/configs"
	"github.com/aru4ka/opinion-hub/internal/app/controllers"
	"github.com/aru4ka/opinion-hub/internal/app/models"
	"github.com/aru4ka/opinion-hub/internal/app/repositories"
	"github.com/aru4ka/opinion-hub/internal/app/services"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

type Server struct {
	config *configs.ServerConfig
	e      *echo.Echo
}

func New(config *configs.ServerConfig) *Server {
	return &Server{config: config, e: echo.New()}
}

func (s *Server) init(db *gorm.DB) {
	userRepo := repositories.NewGormUserRepository(db)
	opinionRepo := repositories.NewGormOpinionRepository(db)

	userService := services.NewUserService(userRepo)
	opinionService := services.NewOpinionService(opinionRepo)

	authController := controllers.NewAuthController(userService)
	userController := controllers.NewUserController(userService, opinionService)

	s.e.Use(middleware.Recover())
	s.e.Use(middleware.CORS())
	s.e.Use(middleware.Logger())

	authGroup := s.e.Group("/auth")
	authGroup.POST("/register", authController.Register())
	authGroup.POST("/login", authController.Login(s.config.TokenSecret, s.config.TokenExpiration))

	userGroup := s.e.Group("/users")
	userGroup.GET("/:username", userController.GetUser())
	userGroup.GET("/:username/opinions", userController.GetOpinions())
}

func (s *Server) Start() error {
	db, err := s.config.Db.GetDB()
	if err != nil {
		return err
	}

	if err := db.AutoMigrate(&models.User{}, &models.Opinion{}, &models.Like{}); err != nil {
		return err
	}

	s.init(db)
	return s.e.Start(":" + s.config.Port)
}
