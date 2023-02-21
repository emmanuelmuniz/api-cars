package repository

import (
	model "api-cars/app/domain/model"
)

type CarModelRepository interface {
	FindAll(cm []*model.CarModel) ([]*model.CarModel, error)
	FindOne(id int) (*model.CarModel, error)
	Create(cm *model.CarModel) (*model.CarModel, error)
	Delete(id int) error
	Update(cm *model.CarModel) (*model.CarModel, error)
}
