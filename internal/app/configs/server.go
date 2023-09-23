package configs

import (
	"os"
	"strconv"

	"github.com/aru4ka/opinion-hub/internal/app/dbs"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBInitializer interface {
	GetDB() (*gorm.DB, error)
}

type ServerConfig struct {
	Port            string
	BaseURL         string
	SMTPEmail       string
	SMTPPassword    string
	TokenSecret     string
	TokenExpiration int

	Db DBInitializer
}

func NewPostgresServerConfig(loadEnv bool) (*ServerConfig, error) {
	if loadEnv {
		godotenv.Load(".env")
	}

	tokenExpiration, err := strconv.Atoi(os.Getenv("TOKEN_EXPIRATION_HRS"))
	if err != nil {
		return nil, err
	}

	return &ServerConfig{
		Port:            os.Getenv("PORT"),
		BaseURL:         os.Getenv("URL"),
		SMTPEmail:       os.Getenv("SMTP_EMAIL"),
		SMTPPassword:    os.Getenv("SMTP_PASSWORD"),
		TokenSecret:     os.Getenv("TOKEN_SECRET"),
		TokenExpiration: tokenExpiration,
		Db: &dbs.PostgresInitializer{
			Host:     os.Getenv("PGHOST"),
			Port:     os.Getenv("PGPORT"),
			DbName:   os.Getenv("PGDATABASE"),
			Username: os.Getenv("PGUSER"),
			Password: os.Getenv("PGPASSWORD"),
			LogLevel: logger.Info,
		},
	}, nil
}
