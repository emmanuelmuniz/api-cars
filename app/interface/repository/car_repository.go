package repository

import (
	"api-cars/app/domain/model"
	"api-cars/app/usecase/repository"
	"errors"

	"gorm.io/gorm"
)

type carRepository struct {
	db *gorm.DB
}

func NewCarRepository(db *gorm.DB) repository.CarRepository {
	return &carRepository{db}
}

func (cr *carRepository) FindAll(c []*model.Car) ([]*model.Car, error) {
	err := cr.db.Find(&c).Error

	if err != nil {
		return nil, err
	}

	return c, nil
}

func (cr *carRepository) FindOne(id string) (*model.Car, error) {
	var c *model.Car

	err := cr.db.First(&c, id).Error

	if c == nil {
		return nil, errors.New("Record with id " + id + "not fond")
	}

	if err != nil {
		return nil, err
	}

	return c, nil
}

func (cr *carRepository) Create(c *model.Car) (*model.Car, error) {
	if err := cr.db.Create(c).Error; err != nil {
		return nil, err
	}

	return c, nil
}

func (cr *carRepository) Delete(id string) {
	var c *model.Car
	cr.db.Delete(&c, id)
}
