package repository

import (
	"api-cars/app/cars-app/car-image/model"

	"gorm.io/gorm"
)

type featureRepository struct {
	db *gorm.DB
}

func NewCarImageRepository(db *gorm.DB) CarImageRepository {
	return &featureRepository{db}
}

func (fr *featureRepository) FindAll(f []*model.CarImage) ([]*model.CarImage, error) {
	err := fr.db.Find(&f).Error

	if err != nil {
		return nil, err
	}

	return f, nil
}

func (fr *featureRepository) FindOne(id int) (*model.CarImage, error) {
	var f *model.CarImage

	err := fr.db.First(&f, id).Error

	if err != nil {
		return nil, err
	}

	return f, nil
}

func (r *featureRepository) FindByIDs(ids []int) ([]*model.CarImage, error) {
	var fts []*model.CarImage

	err := r.db.Find(&fts, ids).Error

	if err != nil {
		return nil, err
	}

	return fts, nil
}

func (fr *featureRepository) Create(f *model.CarImage) (*model.CarImage, error) {
	if err := fr.db.Create(f).Error; err != nil {
		return nil, err
	}

	return f, nil
}

func (fr *featureRepository) Delete(id int) error {
	var f model.CarImage
	err := fr.db.Table("car_features").Where("feature_id = ?", id).Delete(nil).Error

	if err != nil {
		return err
	}

	err = fr.db.Delete(&f, id).Error

	return err
}

func (fr *featureRepository) Update(f *model.CarImage) (*model.CarImage, error) {
	if err := fr.db.Save(f).Error; err != nil {
		return nil, err
	}

	return f, nil
}
