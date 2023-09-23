package dbs

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type PostgresInitializer struct {
	Host     string
	Port     string
	DbName   string
	Username string
	Password string

	LogLevel logger.LogLevel
}

func (p *PostgresInitializer) GetDB() (*gorm.DB, error) {
	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		p.Host,
		p.Port,
		p.Username,
		p.DbName,
		p.Password,
	)

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{
		Logger: logger.Default.LogMode(p.LogLevel),
	})
	if err != nil {
		return nil, err
	}

	if err = db.AutoMigrate(); err != nil {
		return nil, err
	}

	return db, nil
}
