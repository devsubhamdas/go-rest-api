package storage

import (
	"github.com/Subham-Das-98/go-rest-api/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgres(cfg config.Postgres) (*gorm.DB, error) {
	dsn := cfg.DBUrl

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
