package registry

import (
	dbr "api-cars/app/app-common/db"

	"api-cars/app/cars-app/make/controller"
	mp "api-cars/app/cars-app/make/presenter"
	mr "api-cars/app/cars-app/make/repository"
	ms "api-cars/app/cars-app/make/service"
)

func (r *registry) NewMakeController() controller.MakeController {
	makeInteractor := ms.NewMakeService(
		mr.NewMakeRepository(r.db),
		mp.NewMakePresenter(),
		dbr.NewDBRepository(r.db),
	)

	return controller.NewMakeController(makeInteractor)
}
