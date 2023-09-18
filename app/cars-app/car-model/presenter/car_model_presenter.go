package presenter

import (
	"api-cars/app/cars-app/car-model/model"
)

type carModelPresenter struct{}

func NewCarModelPresenter() CarModelPresenter {
	return &carModelPresenter{}
}

func (cmp *carModelPresenter) ResponseCarModels(carModels []*model.CarModel) []*model.CarModel {
	return carModels
}

func (cmp *carModelPresenter) ResponseCarModel(carModel *model.CarModel) *model.CarModel {
	return carModel
}
