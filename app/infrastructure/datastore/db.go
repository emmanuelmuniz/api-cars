package datastore

import (
	bsm "api-cars/app/cars-app/body-style/model"
	cim "api-cars/app/cars-app/car-image/model"
	cmm "api-cars/app/cars-app/car-model/model"
	cm "api-cars/app/cars-app/car/model"
	fm "api-cars/app/cars-app/feature/model"
	mm "api-cars/app/cars-app/make/model"

	"api-cars/app/config"
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

	database.AutoMigrate(&cm.Car{})
	database.AutoMigrate(&mm.Make{})
	database.AutoMigrate(&cmm.CarModel{})
	database.AutoMigrate(&bsm.BodyStyle{})
	database.AutoMigrate(&fm.Feature{})
	database.AutoMigrate(&cim.CarImage{})
	return database
}
