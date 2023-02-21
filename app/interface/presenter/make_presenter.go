package presenter

import (
	"api-cars/app/domain/model"
	"api-cars/app/usecase/presenter"
)

type makePresenter struct{}

func NewMakePresenter() presenter.MakePresenter {
	return &makePresenter{}
}

func (mp *makePresenter) ResponseMakes(makes []*model.Make) []*model.Make {
	return makes
}

func (mp *makePresenter) ResponseMake(make *model.Make) *model.Make {
	return make
}
