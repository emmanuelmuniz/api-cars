package registry

import (
	"api-cars/app/interface/controller"
	ip "api-cars/app/interface/presenter"
	ir "api-cars/app/interface/repository"
	"api-cars/app/usecase/interactor"
)

func (r *registry) NewCarController() controller.CarController {
	carInteractor := interactor.NewCarInteractor(
		ir.NewCarRepository(r.db),
		ip.NewCarPresenter(),
		ir.NewDBRepository(r.db),
	)

	return controller.NewCarController(carInteractor)
}
