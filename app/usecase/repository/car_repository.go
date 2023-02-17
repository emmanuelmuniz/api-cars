package repository

import model "api-cars/app/domain/model"

type CarRepository interface {
	FindAll(c []*model.Car) ([]*model.Car, error)
	Create(c *model.Car) (*model.Car, error)
}
