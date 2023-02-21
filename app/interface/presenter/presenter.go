package presenter

import (
	"api-cars/app/domain/model"
	"api-cars/app/usecase/presenter"
)

type carPresenter struct{}

type makePresenter struct{}

type carModelPresenter struct{}

func NewCarPresenter() presenter.CarPresenter {
	return &carPresenter{}
}

func NewMakePresenter() presenter.MakePresenter {
	return &makePresenter{}
}

func NewCarModelPresenter() presenter.CarModelPresenter {
	return &carModelPresenter{}
}

func (cp *carPresenter) ResponseCars(cars []*model.Car) []*model.Car {
	return cars
}

func (cp *carPresenter) ResponseCar(car *model.Car) *model.Car {
	return car
}

func (mp *makePresenter) ResponseMakes(makes []*model.Make) []*model.Make {
	return makes
}

func (mp *makePresenter) ResponseMake(make *model.Make) *model.Make {
	return make
}

func (cmp *carModelPresenter) ResponseCarModels(carModels []*model.CarModel) []*model.CarModel {
	return carModels
}

func (cmp *carModelPresenter) ResponseCarModel(carModel *model.CarModel) *model.CarModel {
	return carModel
}
