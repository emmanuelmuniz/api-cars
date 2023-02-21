package presenter

import (
	"api-cars/app/domain/model"
	"api-cars/app/usecase/presenter"
)

type carModelPresenter struct{}

func NewCarModelPresenter() presenter.CarModelPresenter {
	return &carModelPresenter{}
}

func (cmp *carModelPresenter) ResponseCarModels(carModels []*model.CarModel) []*model.CarModel {
	return carModels
}

func (cmp *carModelPresenter) ResponseCarModel(carModel *model.CarModel) *model.CarModel {
	return carModel
}
