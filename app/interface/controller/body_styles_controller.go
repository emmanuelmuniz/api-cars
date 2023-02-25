package controller

import (
	"api-cars/app/domain/model"
	"api-cars/app/usecase/interactor"
	"net/http"
)

type bodyStyleController struct {
	bodyStyleInteractor interactor.BodyStyleInteractor
}

type BodyStyleController interface {
	GetBodyStyles(c Context) error
	GetBodyStyle(c Context, id string) error
	CreateBodyStyle(c Context) error
	DeleteBodyStyle(c Context, id string) error
	UpdateBodyStyle(c Context) error
}

func NewBodyStyleController(bodyStyle interactor.BodyStyleInteractor) BodyStyleController {
	return &bodyStyleController{bodyStyle}
}

func (bsc *bodyStyleController) GetBodyStyles(c Context) error {
	var bodyStyle []*model.BodyStyle

	bodyStyle, err := bsc.bodyStyleInteractor.Get(bodyStyle)
	if err != nil {
		return sendErrorCar(c, err)
	}

	return c.JSON(http.StatusOK, bodyStyle)
}

func (bsc *bodyStyleController) GetBodyStyle(c Context, id string) error {
	bodyStyle, err := bsc.bodyStyleInteractor.GetOne(id)

	if err != nil {
		return sendErrorCar(c, err)
	}

	return c.JSON(http.StatusCreated, bodyStyle)
}

func (bsc *bodyStyleController) CreateBodyStyle(c Context) error {
	var params model.BodyStyle

	if err := c.Bind(&params); err != nil {
		return sendErrorCar(c, model.NewAppError("Bad Request. "+err.Error(), http.StatusBadRequest))
	}

	bodyStyle, err := bsc.bodyStyleInteractor.Create(&params)
	if err != nil {
		return sendErrorCar(c, err)
	}

	return c.JSON(http.StatusCreated, bodyStyle)
}

func (bsc *bodyStyleController) DeleteBodyStyle(c Context, id string) error {
	err := bsc.bodyStyleInteractor.Delete(id)

	if err != nil {
		return sendErrorCar(c, err)
	}

	return c.JSON(http.StatusNoContent, nil)
}

func (bsc *bodyStyleController) UpdateBodyStyle(c Context) error {
	var params model.BodyStyle

	if err := c.Bind(&params); err != nil {
		return sendErrorCar(c, model.NewAppError("Bad Request. "+err.Error(), http.StatusBadRequest))
	}

	bodyStyle, err := bsc.bodyStyleInteractor.Update(&params)
	if err != nil {
		return sendErrorCar(c, err)
	}

	return c.JSON(http.StatusCreated, bodyStyle)
}
