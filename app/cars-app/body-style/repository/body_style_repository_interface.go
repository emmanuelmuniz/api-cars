package repository

import (
	"api-cars/app/cars-app/body-style/model"
)

type BodyStyleRepository interface {
	FindAll(b []*model.BodyStyle) ([]*model.BodyStyle, error)
	FindOne(id int) (*model.BodyStyle, error)
	Create(b *model.BodyStyle) (*model.BodyStyle, error)
	Delete(id int) error
	Update(b *model.BodyStyle) (*model.BodyStyle, error)
}
