package registry

import (
	dbr "api-cars/app/app-common/db"

	"api-cars/app/cars-app/body-style/controller"
	"api-cars/app/cars-app/body-style/presenter"
	"api-cars/app/cars-app/body-style/repository"

	service "api-cars/app/cars-app/body-style/service"
)

func (r *registry) NewBodyStyleController() controller.BodyStyleController {
	bodyStyleInteractor := service.NewBodyStyleService(
		repository.NewBodyStyleRepository(r.db),
		presenter.NewBodyStylePresenter(),
		dbr.NewDBRepository(r.db),
	)

	return controller.NewBodyStyleController(bodyStyleInteractor)
}
