package controller

import (
	context "api-cars/app/app-common/context"

	errorCar "api-cars/app/app-controllers/error"
	errorModel "api-cars/app/domain/model"

	service "api-cars/app/cars-app/car-image/service"

	"api-cars/app/cars-app/car-image/model"
	"net/http"
)

type carImageController struct {
	carImageService service.CarImageService
}

type CarImageController interface {
	GetCarImages(c context.Context) error
	GetCarImage(c context.Context, id string) error
	CreateCarImage(c context.Context) error
	DeleteCarImage(c context.Context, id string) error
	UpdateCarImage(c context.Context) error
}

func NewCarImageController(carImage service.CarImageService) CarImageController {
	return &carImageController{carImage}
}

func (fc *carImageController) GetCarImages(c context.Context) error {
	var carImage []*model.CarImage

	carImage, err := fc.carImageService.Get(carImage)
	if err != nil {
		return errorCar.SendErrorCar(c, err)
	}

	return c.JSON(http.StatusOK, carImage)
}

func (fc *carImageController) GetCarImage(c context.Context, id string) error {
	carImage, err := fc.carImageService.GetOne(id)

	if err != nil {
		return errorCar.SendErrorCar(c, err)
	}

	return c.JSON(http.StatusCreated, carImage)
}

func (fc *carImageController) CreateCarImage(c context.Context) error {
	var params model.CarImage

	if err := c.Bind(&params); err != nil {
		return errorCar.SendErrorCar(c, errorModel.NewAppError("Bad Request. "+err.Error(), http.StatusBadRequest))
	}

	carImage, err := fc.carImageService.Create(&params)
	if err != nil {
		return errorCar.SendErrorCar(c, err)
	}

	return c.JSON(http.StatusCreated, carImage)
}

func (fc *carImageController) DeleteCarImage(c context.Context, id string) error {
	err := fc.carImageService.Delete(id)

	if err != nil {
		return errorCar.SendErrorCar(c, err)
	}

	return c.JSON(http.StatusNoContent, nil)
}

func (fc *carImageController) UpdateCarImage(c context.Context) error {
	var params model.CarImage

	if err := c.Bind(&params); err != nil {
		return errorCar.SendErrorCar(c, errorModel.NewAppError("Bad Request. "+err.Error(), http.StatusBadRequest))
	}

	carImage, err := fc.carImageService.Update(&params)
	if err != nil {
		return errorCar.SendErrorCar(c, err)
	}

	return c.JSON(http.StatusCreated, carImage)
}
