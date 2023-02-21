package presenter

import (
	"api-cars/app/domain/model"
	"api-cars/app/usecase/presenter"
)

type carPresenter struct{}

func NewCarPresenter() presenter.CarPresenter {
	return &carPresenter{}
}

func (cp *carPresenter) ResponseCars(cars []*model.Car) []*model.Car {
	return cars
}

func (cp *carPresenter) ResponseCar(car *model.Car) *model.Car {
	return car
}
