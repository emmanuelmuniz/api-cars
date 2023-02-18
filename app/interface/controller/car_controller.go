package controller

import (
	"api-cars/app/domain/model"
	"api-cars/app/usecase/interactor"
	"net/http"

	"github.com/labstack/echo/v4"
)

type carController struct {
	carInteractor interactor.CarInteractor
}

type CarController interface {
	GetCars(c Context) error
	GetCar(c Context, id string) error
	CreateCar(c Context) error
	DeleteCar(c Context, id string) error
}

func NewCarController(car interactor.CarInteractor) CarController {
	return &carController{car}
}

func (cc *carController) GetCars(c Context) error {
	var car []*model.Car

	car, err := cc.carInteractor.Get(car)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, car)
}

func (cc *carController) GetCar(c Context, id string) error {
	car, err := cc.carInteractor.GetOne(id)

	if err != nil {
		return echo.NewHTTPError(404)
	}

	return c.JSON(http.StatusCreated, car)
}

func (cc *carController) CreateCar(c Context) error {
	var params model.Car

	if err := c.Bind(&params); err != nil {
		return err
	}

	car, err := cc.carInteractor.Create(&params)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, car)
}

func (cc *carController) DeleteCar(c Context, id string) error {
	cc.carInteractor.Delete(id)
	return c.JSON(http.StatusNoContent, nil)
}
