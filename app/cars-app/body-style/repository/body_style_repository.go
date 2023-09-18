package repository

import (
	"api-cars/app/cars-app/body-style/model"

	"gorm.io/gorm"
)

type mbodyStyleRepository struct {
	db *gorm.DB
}

func NewBodyStyleRepository(db *gorm.DB) BodyStyleRepository {
	return &mbodyStyleRepository{db}
}

func (bsr *mbodyStyleRepository) FindAll(b []*model.BodyStyle) ([]*model.BodyStyle, error) {
	err := bsr.db.Find(&b).Error

	if err != nil {
		return nil, err
	}

	return b, nil
}

func (bsr *mbodyStyleRepository) FindOne(id int) (*model.BodyStyle, error) {
	var b *model.BodyStyle

	err := bsr.db.First(&b, id).Error

	if err != nil {
		return nil, err
	}

	return b, nil
}

func (bsr *mbodyStyleRepository) Create(b *model.BodyStyle) (*model.BodyStyle, error) {
	if err := bsr.db.Create(b).Error; err != nil {
		return nil, err
	}

	return b, nil
}

func (bsr *mbodyStyleRepository) Delete(id int) error {
	var b *model.BodyStyle
	err := bsr.db.Delete(&b, id).Error
	return err
}

func (bsr *mbodyStyleRepository) Update(b *model.BodyStyle) (*model.BodyStyle, error) {
	if err := bsr.db.Save(b).Error; err != nil {
		return nil, err
	}

	return b, nil
}
