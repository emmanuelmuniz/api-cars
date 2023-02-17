package presenter

import "api-cars/app/domain/model"

type CarPresenter interface {
	ResponseCars(c []*model.Car) []*model.Car
}
