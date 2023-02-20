package repository

import (
	"api-cars/app/domain/model"
	"api-cars/app/usecase/repository"
	"errors"

	"gorm.io/gorm"
)

type makeRepository struct {
	db *gorm.DB
}

func NewMakeRepository(db *gorm.DB) repository.MakeRepository {
	return &makeRepository{db}
}

func (mr *makeRepository) FindAll(m []*model.Make) ([]*model.Make, error) {
	err := mr.db.Find(&m).Error

	if err != nil {
		return nil, err
	}

	return m, nil
}

func (mr *makeRepository) FindOne(id string) (*model.Make, error) {
	var m *model.Make

	err := mr.db.First(&m, id).Error

	if m == nil {
		return nil, errors.New("Record with id " + id + "not fond")
	}

	if err != nil {
		return nil, err
	}

	return m, nil
}

func (mr *makeRepository) Create(m *model.Make) (*model.Make, error) {
	if err := mr.db.Create(m).Error; err != nil {
		return nil, err
	}

	return m, nil
}

func (mr *makeRepository) Delete(id string) error {
	var m *model.Make
	err := mr.db.Delete(&m, id).Error
	return err
}

func (mr *makeRepository) Update(m *model.Make) (*model.Make, error) {
	if err := mr.db.Save(m).Error; err != nil {
		return nil, err
	}

	return m, nil
}
