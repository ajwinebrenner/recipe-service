package server

import (
	"github.com/ajwinebrenner/recipe-service/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectToDatabase(URL string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(URL), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(
		&models.Chef{},
		&models.Meal{},
		&models.Ingredient{},
	)
	return db, nil
}
