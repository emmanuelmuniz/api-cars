package controller

import (
	context "api-cars/app/app-common/context"

	errorCar "api-cars/app/app-controllers/error"
	errorModel "api-cars/app/domain/model"

	service "api-cars/app/cars-app/car/service"

	"api-cars/app/cars-app/car/model"

	"net/http"
)

type carController struct {
	carService service.CarService
}

type CarController interface {
	GetCars(c context.Context) error
	GetCar(c context.Context, id string) error
	CreateCar(c context.Context) error
	DeleteCar(c context.Context, id string) error
	UpdateCar(c context.Context) error
}

func NewCarController(car service.CarService) CarController {
	return &carController{car}
}

func (cc *carController) GetCars(c context.Context) error {
	var car []*model.Car

	car, err := cc.carService.Get(car)

	if err != nil {
		return errorCar.SendErrorCar(c, err)
	}

	return c.JSON(http.StatusOK, car)
}

func (cc *carController) GetCar(c context.Context, id string) error {
	car, err := cc.carService.GetOne(id)

	if err != nil {
		return errorCar.SendErrorCar(c, err)
	}

	return c.JSON(http.StatusCreated, car)
}

func (cc *carController) CreateCar(c context.Context) error {
	var params model.Car

	if err := c.Bind(&params); err != nil {
		return errorCar.SendErrorCar(c, errorModel.NewAppError("Bad Request. "+err.Error(), http.StatusBadRequest))
	}

	car, err := cc.carService.Create(&params)
	if err != nil {
		return errorCar.SendErrorCar(c, err)
	}

	return c.JSON(http.StatusCreated, car)
}

func (cc *carController) DeleteCar(c context.Context, id string) error {
	err := cc.carService.Delete(id)

	if err != nil {
		return errorCar.SendErrorCar(c, err)
	}

	return c.JSON(http.StatusNoContent, nil)
}

func (cc *carController) UpdateCar(c context.Context) error {
	var params model.Car

	if err := c.Bind(&params); err != nil {
		return errorCar.SendErrorCar(c, errorModel.NewAppError("Bad Request. "+err.Error(), http.StatusBadRequest))
	}

	car, err := cc.carService.Update(&params)

	if err != nil {
		return errorCar.SendErrorCar(c, err)
	}

	return c.JSON(http.StatusCreated, car)
}
