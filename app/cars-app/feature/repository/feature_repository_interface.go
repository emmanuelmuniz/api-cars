package repository

import (
	model "api-cars/app/cars-app/feature/model"
)

type FeatureRepository interface {
	FindAll(f []*model.Feature) ([]*model.Feature, error)
	FindOne(id int) (*model.Feature, error)
	FindByIDs(ids []int) ([]*model.Feature, error)
	Create(f *model.Feature) (*model.Feature, error)
	Delete(id int) error
	Update(f *model.Feature) (*model.Feature, error)
}
