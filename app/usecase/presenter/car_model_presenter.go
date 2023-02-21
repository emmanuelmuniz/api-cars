package presenter

import "api-cars/app/domain/model"

type CarModelPresenter interface {
	ResponseCarModels(cm []*model.CarModel) []*model.CarModel
	ResponseCarModel(cm *model.CarModel) *model.CarModel
}
