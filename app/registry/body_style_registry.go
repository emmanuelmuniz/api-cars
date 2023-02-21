package registry

import (
	"api-cars/app/interface/controller"
	ip "api-cars/app/interface/presenter"
	ir "api-cars/app/interface/repository"
	"api-cars/app/usecase/interactor"
)

func (r *registry) NewBodyStyleController() controller.BodyStyleController {
	bodyStyleInteractor := interactor.NewBodyStyleInteractor(
		ir.NewBodyStyleRepository(r.db),
		ip.NewBodyStylePresenter(),
		ir.NewDBRepository(r.db),
	)

	return controller.NewBodyStyleController(bodyStyleInteractor)
}
