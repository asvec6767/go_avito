package database

import (
	"main/internal/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewPostgresConn(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	// авто миграции
	err = db.AutoMigrate(&domain.User{}, &domain.Team{}, &domain.PR{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
