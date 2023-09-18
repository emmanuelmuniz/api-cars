package presenter

import "api-cars/app/cars-app/car/model"

type CarPresenter interface {
	ResponseCars(c []*model.Car) []*model.Car
	ResponseCar(car *model.Car) *model.Car
}
