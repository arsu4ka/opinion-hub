package dbs

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type SqliteInitializer struct {
	DBPath string

	LogLevel logger.LogLevel
}

func (s *SqliteInitializer) GetDB() (*gorm.DB, error) {
	return gorm.Open(sqlite.Open(s.DBPath), &gorm.Config{
		Logger: logger.Default.LogMode(s.LogLevel),
	})
}
