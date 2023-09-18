package presenter

import "api-cars/app/cars-app/body-style/model"

type BodyStylePresenter interface {
	ResponseBodyStyles(b []*model.BodyStyle) []*model.BodyStyle
	ResponseBodyStyle(b *model.BodyStyle) *model.BodyStyle
}
