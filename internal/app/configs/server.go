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

func NewServerConfig(loadEnv bool, dbInitializer DBInitializer) (*ServerConfig, error) {
	if loadEnv {
		err := godotenv.Load(".env")
		if err != nil {
			return nil, err
		}
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
		Db:              dbInitializer,
	}, nil
}

func NewPostgresServerConfig(loadEnv bool) (*ServerConfig, error) {
	if loadEnv {
		err := godotenv.Load(".env")
		if err != nil {
			return nil, err
		}
	}

	dbInitializer := &dbs.PostgresInitializer{
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     os.Getenv("POSTGRES_PORT"),
		DbName:   os.Getenv("POSTGRES_DB"),
		Username: os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		LogLevel: logger.Info,
	}

	return NewServerConfig(loadEnv, dbInitializer)
}

func NewSqliteServerConfig(loadEnv bool) (*ServerConfig, error) {
	if loadEnv {
		err := godotenv.Load(".env")
		if err != nil {
			return nil, err
		}
	}

	dbInitializer := &dbs.SqliteInitializer{
		DBPath:   "DB_PATH",
		LogLevel: logger.Info,
	}

	return NewServerConfig(loadEnv, dbInitializer)
}
