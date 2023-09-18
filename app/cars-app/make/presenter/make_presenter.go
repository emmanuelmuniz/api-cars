package presenter

import (
	"api-cars/app/cars-app/make/model"
)

type makePresenter struct{}

func NewMakePresenter() MakePresenter {
	return &makePresenter{}
}

func (mp *makePresenter) ResponseMakes(makes []*model.Make) []*model.Make {
	return makes
}

func (mp *makePresenter) ResponseMake(make *model.Make) *model.Make {
	return make
}
