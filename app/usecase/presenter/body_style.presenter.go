package presenter

import "api-cars/app/domain/model"

type BodyStylePresenter interface {
	ResponseBodyStyles(b []*model.BodyStyle) []*model.BodyStyle
	ResponseBodyStyle(b *model.BodyStyle) *model.BodyStyle
}
