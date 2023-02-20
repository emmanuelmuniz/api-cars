package registry

import (
	"api-cars/app/interface/controller"
	ip "api-cars/app/interface/presenter"
	ir "api-cars/app/interface/repository"
	"api-cars/app/usecase/interactor"
)

func (r *registry) NewCarModelController() controller.CarModelController {
	carModelInteractor := interactor.NewCarModelInteractor(
		ir.NewCarModelRepository(r.db),
		ip.NewCarModelPresenter(),
		ir.NewDBRepository(r.db),
	)

	return controller.NewCarModelController(carModelInteractor)
}
