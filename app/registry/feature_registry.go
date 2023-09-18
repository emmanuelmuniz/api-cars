package registry

import (
	dbr "api-cars/app/app-common/db"

	"api-cars/app/cars-app/feature/controller"
	"api-cars/app/cars-app/feature/presenter"
	"api-cars/app/cars-app/feature/repository"

	service "api-cars/app/cars-app/feature/service"
)

func (r *registry) NewFeatureController() controller.FeatureController {
	featureInteractor := service.NewFeatureService(
		repository.NewFeatureRepository(r.db),
		presenter.NewFeaturePresenter(),
		dbr.NewDBRepository(r.db),
	)

	return controller.NewFeatureController(featureInteractor)
}
