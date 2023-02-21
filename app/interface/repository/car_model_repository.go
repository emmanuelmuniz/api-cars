package repository

import (
	"api-cars/app/domain/model"
	"api-cars/app/usecase/repository"

	"gorm.io/gorm"
)

type carModelRepository struct {
	db *gorm.DB
}

func NewCarModelRepository(db *gorm.DB) repository.CarModelRepository {
	return &carModelRepository{db}
}

func (mr *carModelRepository) FindAll(cm []*model.CarModel) ([]*model.CarModel, error) {
	err := mr.db.Preload("Make").Find(&cm).Error

	if err != nil {
		return nil, err
	}

	return cm, nil
}

func (mr *carModelRepository) FindOne(id int) (*model.CarModel, error) {
	var cm *model.CarModel

	err := mr.db.Preload("Make").First(&cm, id).Error

	if err != nil {
		return nil, err
	}

	return cm, nil
}

func (mr *carModelRepository) Create(cm *model.CarModel) (*model.CarModel, error) {
	if err := mr.db.Create(cm).Error; err != nil {
		return nil, err
	}

	return cm, nil
}

func (mr *carModelRepository) Delete(id int) error {
	var cm *model.CarModel
	err := mr.db.Delete(&cm, id).Error
	return err
}

func (mr *carModelRepository) Update(cm *model.CarModel) (*model.CarModel, error) {
	if err := mr.db.Save(cm).Error; err != nil {
		return nil, err
	}

	return cm, nil
}
