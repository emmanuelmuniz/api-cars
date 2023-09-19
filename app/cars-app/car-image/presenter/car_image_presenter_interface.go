package presenter

import "api-cars/app/cars-app/car-image/model"

type CarImagePresenter interface {
	ResponseCarImages(f []*model.CarImage) []*model.CarImage
	ResponseCarImage(f *model.CarImage) *model.CarImage
}
