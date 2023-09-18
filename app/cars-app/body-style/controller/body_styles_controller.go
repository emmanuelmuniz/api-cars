package controller

import (
	context "api-cars/app/app-common/context"

	errorCar "api-cars/app/app-controllers/error"
	errorModel "api-cars/app/domain/model"

	service "api-cars/app/cars-app/body-style/service"

	"api-cars/app/cars-app/body-style/model"

	"net/http"
)

type bodyStyleController struct {
	bodyStyleService service.BodyStyleService
}

type BodyStyleController interface {
	GetBodyStyles(c context.Context) error
	GetBodyStyle(c context.Context, id string) error
	CreateBodyStyle(c context.Context) error
	DeleteBodyStyle(c context.Context, id string) error
	UpdateBodyStyle(c context.Context) error
}

func NewBodyStyleController(bodyStyle service.BodyStyleService) BodyStyleController {
	return &bodyStyleController{bodyStyle}
}

func (bsc *bodyStyleController) GetBodyStyles(c context.Context) error {
	var bodyStyle []*model.BodyStyle

	bodyStyle, err := bsc.bodyStyleService.Get(bodyStyle)

	if err != nil {
		return errorCar.SendErrorCar(c, err)
	}

	return c.JSON(http.StatusOK, bodyStyle)
}

func (bsc *bodyStyleController) GetBodyStyle(c context.Context, id string) error {
	bodyStyle, err := bsc.bodyStyleService.GetOne(id)

	if err != nil {
		return errorCar.SendErrorCar(c, err)
	}

	return c.JSON(http.StatusCreated, bodyStyle)
}

func (bsc *bodyStyleController) CreateBodyStyle(c context.Context) error {
	var params model.BodyStyle

	if err := c.Bind(&params); err != nil {
		return errorCar.SendErrorCar(c, errorModel.NewAppError("Bad Request. "+err.Error(), http.StatusBadRequest))
	}

	bodyStyle, err := bsc.bodyStyleService.Create(&params)
	if err != nil {
		return errorCar.SendErrorCar(c, err)
	}

	return c.JSON(http.StatusCreated, bodyStyle)
}

func (bsc *bodyStyleController) DeleteBodyStyle(c context.Context, id string) error {
	err := bsc.bodyStyleService.Delete(id)

	if err != nil {
		return errorCar.SendErrorCar(c, err)
	}

	return c.JSON(http.StatusNoContent, nil)
}

func (bsc *bodyStyleController) UpdateBodyStyle(c context.Context) error {
	var params model.BodyStyle

	if err := c.Bind(&params); err != nil {
		return errorCar.SendErrorCar(c, errorModel.NewAppError("Bad Request. "+err.Error(), http.StatusBadRequest))
	}

	bodyStyle, err := bsc.bodyStyleService.Update(&params)
	if err != nil {
		return errorCar.SendErrorCar(c, err)
	}

	return c.JSON(http.StatusCreated, bodyStyle)
}
