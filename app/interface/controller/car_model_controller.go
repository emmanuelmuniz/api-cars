package controller

import (
	"api-cars/app/domain/model"
	"api-cars/app/usecase/interactor"
	"net/http"

	"github.com/labstack/echo/v4"
)

type carModelController struct {
	carModelInteractor interactor.CarModelInteractor
}

type CarModelController interface {
	GetCarModels(c Context) error
	GetCarModel(c Context, id string) error
	CreateCarModel(c Context) error
	DeleteCarModel(c Context, id string) error
	UpdateCarModel(c Context) error
}

func NewCarModelController(carModel interactor.CarModelInteractor) CarModelController {
	return &carModelController{carModel}
}

func (cmc *carModelController) GetCarModels(c Context) error {
	var carModel []*model.CarModel

	carModel, err := cmc.carModelInteractor.Get(carModel)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, carModel)
}

func (cmc *carModelController) GetCarModel(c Context, id string) error {
	carModel, err := cmc.carModelInteractor.GetOne(id)

	if err != nil {
		return echo.NewHTTPError(404)
	}

	return c.JSON(http.StatusCreated, carModel)
}

func (cmc *carModelController) CreateCarModel(c Context) error {
	var params model.CarModel

	if err := c.Bind(&params); err != nil {
		return err
	}

	carModel, err := cmc.carModelInteractor.Create(&params)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, carModel)
}

func (cmc *carModelController) DeleteCarModel(c Context, id string) error {
	err := cmc.carModelInteractor.Delete(id)

	if err != nil {
		return echo.NewHTTPError(404)
	}

	return c.JSON(http.StatusNoContent, nil)
}

func (cmc *carModelController) UpdateCarModel(c Context) error {
	var params model.CarModel

	if err := c.Bind(&params); err != nil {
		return err
	}

	carModel, err := cmc.carModelInteractor.Update(&params)
	if err != nil {
		return echo.NewHTTPError(404)
	}

	return c.JSON(http.StatusCreated, carModel)
}
