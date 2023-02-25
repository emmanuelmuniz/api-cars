package registry

import (
	"api-cars/app/interface/controller"
	ip "api-cars/app/interface/presenter"
	ir "api-cars/app/interface/repository"
	"api-cars/app/usecase/interactor"
)

func (r *registry) NewFeatureController() controller.FeatureController {
	featureInteractor := interactor.NewFeatureInteractor(
		ir.NewFeatureRepository(r.db),
		ip.NewFeaturePresenter(),
		ir.NewDBRepository(r.db),
	)

	return controller.NewFeatureController(featureInteractor)
}
