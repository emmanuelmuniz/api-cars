package repository

import (
	model "api-cars/app/domain/model"
)

type CarRepository interface {
	FindAll(c []*model.Car) ([]*model.Car, error)
	FindOne(id string) (*model.Car, error)
	Create(c *model.Car) (*model.Car, error)
	Delete(id string)
	Update(c *model.Car) (*model.Car, error)
}
