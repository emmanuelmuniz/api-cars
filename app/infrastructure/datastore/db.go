package datastore

import (
	"api-cars/app/config"
	"api-cars/app/domain/model"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func NewDB() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.C.Database.Addr,
		config.C.Database.User,
		config.C.Database.Password,
		config.C.Database.DBName,
		config.C.Database.Port)
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database")
	}

	database.AutoMigrate(&model.Car{})
	database.AutoMigrate(&model.Make{})
	database.AutoMigrate(&model.CarModel{})

	return database
}
