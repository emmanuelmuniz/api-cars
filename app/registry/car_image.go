package registry

import (
	dbr "api-cars/app/app-common/db"

	"api-cars/app/cars-app/car-image/controller"
	"api-cars/app/cars-app/car-image/presenter"
	"api-cars/app/cars-app/car-image/repository"

	service "api-cars/app/cars-app/car-image/service"
)

func (r *registry) NewCarImageController() controller.CarImageController {
	carImageInteractor := service.NewCarImageService(
		repository.NewCarImageRepository(r.db),
		presenter.NewCarImagePresenter(),
		dbr.NewDBRepository(r.db),
	)

	return controller.NewCarImageController(carImageInteractor)
}
