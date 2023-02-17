package presenter

import (
	"api-cars/app/domain/model"
	"api-cars/app/usecase/presenter"
)

type carPresenter struct{}

func NewCarPresenter() presenter.CarPresenter {
	return &carPresenter{}
}

func (up *carPresenter) ResponseCars(us []*model.Car) []*model.Car {
	for _, u := range us {
		u.Make = "Car Make: " + u.Make
	}
	return us
}
