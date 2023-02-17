package controller

import (
	"api-cars/app/domain/model"
	"api-cars/app/usecase/interactor"
	"net/http"
)

type carController struct {
	carInteractor interactor.CarInteractor
}

type CarController interface {
	GetCars(c Context) error
	CreateCar(c Context) error
}

func NewCarController(car interactor.CarInteractor) CarController {
	return &carController{car}
}

func (cc *carController) GetCars(c Context) error {
	var u []*model.Car

	u, err := cc.carInteractor.Get(u)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, u)
}

func (cc *carController) CreateCar(c Context) error {
	var params model.Car

	if err := c.Bind(&params); err != nil {
		return err
	}

	u, err := cc.carInteractor.Create(&params)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, u)
}
