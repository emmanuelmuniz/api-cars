package repository

import (
	model "api-cars/app/domain/model"
)

type MakeRepository interface {
	FindAll(m []*model.Make) ([]*model.Make, error)
	FindOne(id string) (*model.Make, error)
	Create(m *model.Make) (*model.Make, error)
	Delete(id string) error
	Update(m *model.Make) (*model.Make, error)
}
