package repository

import (
	"api-cars/app/cars-app/car/model"

	"gorm.io/gorm"
)

type carRepository struct {
	db *gorm.DB
}

func NewCarRepository(db *gorm.DB) CarRepository {
	return &carRepository{db}
}

func (cr *carRepository) FindAll(c []*model.Car) ([]*model.Car, error) {
	err := cr.db.Preload("Make").Preload("CarModel").Preload("CarModel.Make").Preload("BodyStyle").Preload("Features").Find(&c).Error

	if err != nil {
		return nil, err
	}

	return c, nil
}

func (cr *carRepository) FindOne(id int) (*model.Car, error) {
	var c *model.Car
	err := cr.db.Preload("Make").Preload("CarModel").Preload("CarModel.Make").Preload("BodyStyle").Preload("Features").First(&c, id).Error

	if err != nil {
		return nil, err
	}

	return c, nil
}

func (cr *carRepository) Create(c *model.Car) (*model.Car, error) {

	if err := cr.db.Create(&c).Error; err != nil {
		return nil, err
	}

	return c, nil
}

func (cr *carRepository) Delete(id int) error {
	var c *model.Car
	err := cr.db.Delete(&c, id).Error
	return err
}

func (cr *carRepository) Update(c *model.Car) (*model.Car, error) {
	if err := cr.db.Save(c).Error; err != nil {
		return nil, err
	}

	return c, nil
}
