package presenter

import "api-cars/app/cars-app/car-model/model"

type CarModelPresenter interface {
	ResponseCarModels(cm []*model.CarModel) []*model.CarModel
	ResponseCarModel(cm *model.CarModel) *model.CarModel
}
