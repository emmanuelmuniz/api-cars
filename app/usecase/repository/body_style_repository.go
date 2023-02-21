package repository

import (
	model "api-cars/app/domain/model"
)

type BodyStyleRepository interface {
	FindAll(b []*model.BodyStyle) ([]*model.BodyStyle, error)
	FindOne(id int) (*model.BodyStyle, error)
	Create(b *model.BodyStyle) (*model.BodyStyle, error)
	Delete(id int) error
	Update(b *model.BodyStyle) (*model.BodyStyle, error)
}
