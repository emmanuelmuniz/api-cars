package repository

import (
	model "api-cars/app/cars-app/car-image/model"
)

type CarImageRepository interface {
	FindAll(f []*model.CarImage) ([]*model.CarImage, error)
	FindOne(id int) (*model.CarImage, error)
	FindByIDs(ids []int) ([]*model.CarImage, error)
	Create(f *model.CarImage) (*model.CarImage, error)
	Delete(id int) error
	Update(f *model.CarImage) (*model.CarImage, error)
}
