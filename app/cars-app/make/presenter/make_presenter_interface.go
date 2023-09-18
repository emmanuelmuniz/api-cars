package presenter

import "api-cars/app/cars-app/make/model"

type MakePresenter interface {
	ResponseMakes(m []*model.Make) []*model.Make
	ResponseMake(m *model.Make) *model.Make
}
