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
	GetCar(c Context, id string) error
	CreateCar(c Context) error
	DeleteCar(c Context, id string) error
	UpdateCar(c Context) error
}

func NewCarController(car interactor.CarInteractor) CarController {
	return &carController{car}
}

func (cc *carController) GetCars(c Context) error {
	var car []*model.Car

	car, err := cc.carInteractor.Get(car)

	if err != nil {
		return sendError(c, err)
	}

	return c.JSON(http.StatusOK, car)
}

func (cc *carController) GetCar(c Context, id string) error {
	car, err := cc.carInteractor.GetOne(id)

	if err != nil {
		return sendError(c, err)
	}

	return c.JSON(http.StatusCreated, car)
}

func (cc *carController) CreateCar(c Context) error {
	var params model.Car

	if err := c.Bind(&params); err != nil {
		return sendError(c, err)
	}

	car, err := cc.carInteractor.Create(&params)
	if err != nil {
		return sendError(c, err)
	}

	return c.JSON(http.StatusCreated, car)
}

func (cc *carController) DeleteCar(c Context, id string) error {
	err := cc.carInteractor.Delete(id)

	if err != nil {
		return sendError(c, err)
	}

	return c.JSON(http.StatusNoContent, nil)
}

func (cc *carController) UpdateCar(c Context) error {
	var params model.Car

	if err := c.Bind(&params); err != nil {
		return sendError(c, err)
	}

	car, err := cc.carInteractor.Update(&params)

	if err != nil {
		return sendError(c, err)
	}

	return c.JSON(http.StatusCreated, car)
}

func sendError(c Context, err error) error {
	e, ok := err.(*model.AppError)

	if e.Code != http.StatusInternalServerError && ok {
		return c.JSON(e.Code, map[string]string{
			"message": e.Message,
		})
	}

	return c.JSON(http.StatusInternalServerError, map[string]string{
		"error": "Internal Server Error",
	})
}
