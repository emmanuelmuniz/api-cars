package presenter

import (
	"api-cars/app/cars-app/car-image/model"
)

type carImagePresenter struct{}

func NewCarImagePresenter() CarImagePresenter {
	return &carImagePresenter{}
}

func (fp *carImagePresenter) ResponseCarImages(carImages []*model.CarImage) []*model.CarImage {
	return carImages
}

func (fp *carImagePresenter) ResponseCarImage(carImage *model.CarImage) *model.CarImage {
	return carImage
}
