package presenter

import (
	"api-cars/app/cars-app/body-style/model"
)

type bodyStylePresenter struct{}

func NewBodyStylePresenter() BodyStylePresenter {
	return &bodyStylePresenter{}
}

func (bsp *bodyStylePresenter) ResponseBodyStyles(bodyStyles []*model.BodyStyle) []*model.BodyStyle {
	return bodyStyles
}

func (bsp *bodyStylePresenter) ResponseBodyStyle(bodyStyle *model.BodyStyle) *model.BodyStyle {
	return bodyStyle
}
