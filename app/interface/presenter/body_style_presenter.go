package presenter

import (
	"api-cars/app/domain/model"
	"api-cars/app/usecase/presenter"
)

type bodyStylePresenter struct{}

func NewBodyStylePresenter() presenter.BodyStylePresenter {
	return &bodyStylePresenter{}
}

func (bsp *bodyStylePresenter) ResponseBodyStyles(bodyStyles []*model.BodyStyle) []*model.BodyStyle {
	return bodyStyles
}

func (bsp *bodyStylePresenter) ResponseBodyStyle(bodyStyle *model.BodyStyle) *model.BodyStyle {
	return bodyStyle
}
