package controller

import (
	"api-cars/app/domain/model"
	"api-cars/app/usecase/interactor"
	"net/http"

	"github.com/labstack/echo/v4"
)

type makeController struct {
	makeInteractor interactor.MakeInteractor
}

type MakeController interface {
	GetMakes(c Context) error
	GetMake(c Context, id string) error
	CreateMake(c Context) error
	DeleteMake(c Context, id string) error
	UpdateMake(c Context) error
}

func NewMakeController(make interactor.MakeInteractor) MakeController {
	return &makeController{make}
}

func (mc *makeController) GetMakes(c Context) error {
	var make []*model.Make

	make, err := mc.makeInteractor.Get(make)
	if err != nil {
		return sendErrorCar(c, err)
	}

	return c.JSON(http.StatusOK, make)
}

func (mc *makeController) GetMake(c Context, id string) error {
	make, err := mc.makeInteractor.GetOne(id)

	if err != nil {
		return echo.NewHTTPError(404)
	}

	return c.JSON(http.StatusCreated, make)
}

func (mc *makeController) CreateMake(c Context) error {
	var params model.Make

	if err := c.Bind(&params); err != nil {
		return sendErrorCar(c, model.NewAppError("Bad Request. "+err.Error(), http.StatusBadRequest))
	}

	make, err := mc.makeInteractor.Create(&params)
	if err != nil {
		return sendErrorCar(c, err)
	}

	return c.JSON(http.StatusCreated, make)
}

func (mc *makeController) DeleteMake(c Context, id string) error {
	err := mc.makeInteractor.Delete(id)

	if err != nil {
		return sendErrorCar(c, err)
	}

	return c.JSON(http.StatusNoContent, nil)
}

func (mc *makeController) UpdateMake(c Context) error {
	var params model.Make

	if err := c.Bind(&params); err != nil {
		return sendErrorCar(c, model.NewAppError("Bad Request. "+err.Error(), http.StatusBadRequest))
	}

	make, err := mc.makeInteractor.Update(&params)
	if err != nil {
		return sendErrorCar(c, err)
	}

	return c.JSON(http.StatusCreated, make)
}
