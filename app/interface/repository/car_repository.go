package repository

import (
	"api-cars/app/domain/model"
	"api-cars/app/usecase/repository"

	"gorm.io/gorm"
)

type carRepository struct {
	db *gorm.DB
}

func NewCarRepository(db *gorm.DB) repository.CarRepository {
	return &carRepository{db}
}

func (ur *carRepository) FindAll(u []*model.Car) ([]*model.Car, error) {
	err := ur.db.Find(&u).Error

	if err != nil {
		return nil, err
	}

	return u, nil
}

func (ur *carRepository) Create(u *model.Car) (*model.Car, error) {
	if err := ur.db.Create(u).Error; err != nil {
		return nil, err
	}

	return u, nil
}
