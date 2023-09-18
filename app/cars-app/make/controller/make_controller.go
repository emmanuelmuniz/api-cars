package controller

import (
	context "api-cars/app/app-common/context"

	errorCar "api-cars/app/app-controllers/error"
	errorModel "api-cars/app/domain/model"

	service "api-cars/app/cars-app/make/service"

	"api-cars/app/cars-app/make/model"
	"net/http"
)

type makeController struct {
	makeService service.MakeService
}

type MakeController interface {
	GetMakes(c context.Context) error
	GetMake(c context.Context, id string) error
	CreateMake(c context.Context) error
	DeleteMake(c context.Context, id string) error
	UpdateMake(c context.Context) error
}

func NewMakeController(make service.MakeService) MakeController {
	return &makeController{make}
}

func (mc *makeController) GetMakes(c context.Context) error {
	var make []*model.Make

	make, err := mc.makeService.Get(make)
	if err != nil {
		return errorCar.SendErrorCar(c, err)
	}

	return c.JSON(http.StatusOK, make)
}

func (mc *makeController) GetMake(c context.Context, id string) error {
	make, err := mc.makeService.GetOne(id)

	if err != nil {
		return errorCar.SendErrorCar(c, err)
	}

	return c.JSON(http.StatusOK, make)
}

func (mc *makeController) CreateMake(c context.Context) error {
	var params model.Make

	if err := c.Bind(&params); err != nil {
		return errorCar.SendErrorCar(c, errorModel.NewAppError("Bad Request. "+err.Error(), http.StatusBadRequest))
	}

	make, err := mc.makeService.Create(&params)
	if err != nil {
		return errorCar.SendErrorCar(c, err)
	}

	return c.JSON(http.StatusCreated, make)
}

func (mc *makeController) DeleteMake(c context.Context, id string) error {
	err := mc.makeService.Delete(id)

	if err != nil {
		return errorCar.SendErrorCar(c, err)
	}

	return c.JSON(http.StatusNoContent, nil)
}

func (mc *makeController) UpdateMake(c context.Context) error {
	var params model.Make

	if err := c.Bind(&params); err != nil {
		return errorCar.SendErrorCar(c, errorModel.NewAppError("Bad Request. "+err.Error(), http.StatusBadRequest))
	}

	make, err := mc.makeService.Update(&params)
	if err != nil {
		return errorCar.SendErrorCar(c, err)
	}

	return c.JSON(http.StatusAccepted, make)
}
