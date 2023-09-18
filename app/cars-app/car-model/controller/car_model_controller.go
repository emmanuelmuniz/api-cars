package controller

import (
	context "api-cars/app/app-common/context"

	errorCar "api-cars/app/app-controllers/error"
	errorModel "api-cars/app/domain/model"

	service "api-cars/app/cars-app/car-model/service"

	"api-cars/app/cars-app/car-model/model"

	"net/http"
)

type carModelController struct {
	carModelService service.CarModelService
}

type CarModelController interface {
	GetCarModels(c context.Context) error
	GetCarModel(c context.Context, id string) error
	CreateCarModel(c context.Context) error
	DeleteCarModel(c context.Context, id string) error
	UpdateCarModel(c context.Context) error
}

func NewCarModelController(carModel service.CarModelService) CarModelController {
	return &carModelController{carModel}
}

func (cmc *carModelController) GetCarModels(c context.Context) error {
	var carModel []*model.CarModel

	carModel, err := cmc.carModelService.Get(carModel)

	if err != nil {
		return errorCar.SendErrorCar(c, err)
	}

	return c.JSON(http.StatusOK, carModel)
}

func (cmc *carModelController) GetCarModel(c context.Context, id string) error {
	carModel, err := cmc.carModelService.GetOne(id)

	if err != nil {
		return errorCar.SendErrorCar(c, err)
	}

	return c.JSON(http.StatusCreated, carModel)
}

func (cmc *carModelController) CreateCarModel(c context.Context) error {
	var params model.CarModel

	if err := c.Bind(&params); err != nil {
		return errorCar.SendErrorCar(c, errorModel.NewAppError("Bad Request. "+err.Error(), http.StatusBadRequest))
	}

	carModel, err := cmc.carModelService.Create(&params)
	if err != nil {
		return errorCar.SendErrorCar(c, err)
	}

	return c.JSON(http.StatusCreated, carModel)
}

func (cmc *carModelController) DeleteCarModel(c context.Context, id string) error {
	err := cmc.carModelService.Delete(id)

	if err != nil {
		return errorCar.SendErrorCar(c, err)
	}

	return c.JSON(http.StatusNoContent, nil)
}

func (cmc *carModelController) UpdateCarModel(c context.Context) error {
	var params model.CarModel

	if err := c.Bind(&params); err != nil {
		return errorCar.SendErrorCar(c, err)
	}

	carModel, err := cmc.carModelService.Update(&params)
	if err != nil {
		return errorCar.SendErrorCar(c, errorModel.NewAppError("Bad Request. "+err.Error(), http.StatusBadRequest))
	}
	return c.JSON(http.StatusCreated, carModel)
}
