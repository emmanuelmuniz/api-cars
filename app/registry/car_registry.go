package registry

import (
	dbr "api-cars/app/app-common/db"
	bsr "api-cars/app/cars-app/body-style/repository"
	cmr "api-cars/app/cars-app/car-model/repository"
	"api-cars/app/cars-app/car/controller"
	cp "api-cars/app/cars-app/car/presenter"
	cr "api-cars/app/cars-app/car/repository"
	cs "api-cars/app/cars-app/car/service"
	fr "api-cars/app/cars-app/feature/repository"
	mr "api-cars/app/cars-app/make/repository"
)

func (r *registry) NewCarController() controller.CarController {
	carService := cs.NewCarService(
		cr.NewCarRepository(r.db),
		cp.NewCarPresenter(),
		mr.NewMakeRepository(r.db),
		cmr.NewCarModelRepository(r.db),
		bsr.NewBodyStyleRepository(r.db),
		fr.NewFeatureRepository(r.db),
		dbr.NewDBRepository(r.db),
	)

	return controller.NewCarController(carService)
}
