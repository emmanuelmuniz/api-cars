package presenter

import "api-cars/app/domain/model"

type MakePresenter interface {
	ResponseMakes(m []*model.Make) []*model.Make
	ResponseMake(m *model.Make) *model.Make
}
