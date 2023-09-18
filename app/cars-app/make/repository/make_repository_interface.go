package repository

import (
	model "api-cars/app/cars-app/make/model"
)

type MakeRepository interface {
	FindAll(m []*model.Make) ([]*model.Make, error)
	FindOne(id int) (*model.Make, error)
	Create(m *model.Make) (*model.Make, error)
	Delete(id int) error
	Update(m *model.Make) (*model.Make, error)
}
