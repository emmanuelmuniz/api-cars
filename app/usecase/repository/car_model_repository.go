package repository

import (
	model "api-cars/app/domain/model"
)

type CarModelRepository interface {
	FindAll(cm []*model.CarModel) ([]*model.CarModel, error)
	FindOne(id string) (*model.CarModel, error)
	Create(cm *model.CarModel) (*model.CarModel, error)
	Delete(id string) error
	Update(cm *model.CarModel) (*model.CarModel, error)
}
