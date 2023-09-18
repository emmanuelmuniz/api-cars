package presenter

import (
	"api-cars/app/cars-app/car/model"
)

type carPresenter struct{}

func NewCarPresenter() CarPresenter {
	return &carPresenter{}
}

func (cp *carPresenter) ResponseCars(cars []*model.Car) []*model.Car {
	return cars
}

func (cp *carPresenter) ResponseCar(car *model.Car) *model.Car {
	return car
}
