package repository

import (
	"github.com/emmanuelmuniz/api-cars/db"
	"github.com/emmanuelmuniz/api-cars/main/cars/models"
)

func GetCars() []models.Car {
	var cars []models.Car
	db.DB.Find(&cars)

	return cars
}

func CreateCar(car models.Car) models.Car {

	return
}
