package database

import (
	"doramaPro/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=bagzhan dbname=doramaPro port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.User{}, &models.Genre{}, &models.Drama{}) // ✅ Миграция

	DB = db
	return db, nil
}
