package controller

import (
	context "api-cars/app/app-common/context"

	errorCar "api-cars/app/app-controllers/error"
	errorModel "api-cars/app/domain/model"

	service "api-cars/app/cars-app/feature/service"

	"api-cars/app/cars-app/feature/model"
	"net/http"
)

type featureController struct {
	featureService service.FeatureService
}

type FeatureController interface {
	GetFeatures(c context.Context) error
	GetFeature(c context.Context, id string) error
	CreateFeature(c context.Context) error
	DeleteFeature(c context.Context, id string) error
	UpdateFeature(c context.Context) error
}

func NewFeatureController(feature service.FeatureService) FeatureController {
	return &featureController{feature}
}

func (fc *featureController) GetFeatures(c context.Context) error {
	var feature []*model.Feature

	feature, err := fc.featureService.Get(feature)
	if err != nil {
		return errorCar.SendErrorCar(c, err)
	}

	return c.JSON(http.StatusOK, feature)
}

func (fc *featureController) GetFeature(c context.Context, id string) error {
	feature, err := fc.featureService.GetOne(id)

	if err != nil {
		return errorCar.SendErrorCar(c, err)
	}

	return c.JSON(http.StatusCreated, feature)
}

func (fc *featureController) CreateFeature(c context.Context) error {
	var params model.Feature

	if err := c.Bind(&params); err != nil {
		return errorCar.SendErrorCar(c, errorModel.NewAppError("Bad Request. "+err.Error(), http.StatusBadRequest))
	}

	feature, err := fc.featureService.Create(&params)
	if err != nil {
		return errorCar.SendErrorCar(c, err)
	}

	return c.JSON(http.StatusCreated, feature)
}

func (fc *featureController) DeleteFeature(c context.Context, id string) error {
	err := fc.featureService.Delete(id)

	if err != nil {
		return errorCar.SendErrorCar(c, err)
	}

	return c.JSON(http.StatusNoContent, nil)
}

func (fc *featureController) UpdateFeature(c context.Context) error {
	var params model.Feature

	if err := c.Bind(&params); err != nil {
		return errorCar.SendErrorCar(c, errorModel.NewAppError("Bad Request. "+err.Error(), http.StatusBadRequest))
	}

	feature, err := fc.featureService.Update(&params)
	if err != nil {
		return errorCar.SendErrorCar(c, err)
	}

	return c.JSON(http.StatusCreated, feature)
}
