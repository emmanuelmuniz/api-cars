package service

import (
	"github.com/emmanuelmuniz/api-cars/main/cars/models"
	"github.com/emmanuelmuniz/api-cars/main/cars/repository"
)

func GetCars() []models.Car {
	return repository.GetCars()
}

func CreateCar(car Car) Car {
	return repository.CreateCar()
}
