package registry

import (
	"api-cars/app/interface/controller"
	ip "api-cars/app/interface/presenter"
	ir "api-cars/app/interface/repository"
	"api-cars/app/usecase/interactor"
)

func (r *registry) NewMakeController() controller.MakeController {
	makeInteractor := interactor.NewMakeInteractor(
		ir.NewMakeRepository(r.db),
		ip.NewMakePresenter(),
		ir.NewDBRepository(r.db),
	)

	return controller.NewMakeController(makeInteractor)
}
