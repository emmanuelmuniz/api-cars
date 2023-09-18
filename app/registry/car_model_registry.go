package registry

import (
	dbr "api-cars/app/app-common/db"
	"api-cars/app/cars-app/car-model/controller"
	cmp "api-cars/app/cars-app/car-model/presenter"
	cmr "api-cars/app/cars-app/car-model/repository"
	cms "api-cars/app/cars-app/car-model/service"
	mr "api-cars/app/cars-app/make/repository"
)

func (r *registry) NewCarModelController() controller.CarModelController {
	carModelInteractor := cms.NewCarModelService(
		cmr.NewCarModelRepository(r.db),
		mr.NewMakeRepository(r.db),
		cmp.NewCarModelPresenter(),
		dbr.NewDBRepository(r.db),
	)

	return controller.NewCarModelController(carModelInteractor)
}
