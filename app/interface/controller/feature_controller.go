package controller

import (
	"api-cars/app/domain/model"
	"api-cars/app/usecase/interactor"
	"net/http"
)

type featureController struct {
	featureInteractor interactor.FeatureInteractor
}

type FeatureController interface {
	GetFeatures(c Context) error
	GetFeature(c Context, id string) error
	CreateFeature(c Context) error
	DeleteFeature(c Context, id string) error
	UpdateFeature(c Context) error
}

func NewFeatureController(feature interactor.FeatureInteractor) FeatureController {
	return &featureController{feature}
}

func (fc *featureController) GetFeatures(c Context) error {
	var feature []*model.Feature

	feature, err := fc.featureInteractor.Get(feature)
	if err != nil {
		return sendErrorCar(c, err)
	}

	return c.JSON(http.StatusOK, feature)
}

func (fc *featureController) GetFeature(c Context, id string) error {
	feature, err := fc.featureInteractor.GetOne(id)

	if err != nil {
		return sendErrorCar(c, err)
	}

	return c.JSON(http.StatusCreated, feature)
}

func (fc *featureController) CreateFeature(c Context) error {
	var params model.Feature

	if err := c.Bind(&params); err != nil {
		return sendErrorCar(c, model.NewAppError("Bad Request. "+err.Error(), http.StatusBadRequest))
	}

	feature, err := fc.featureInteractor.Create(&params)
	if err != nil {
		return sendErrorCar(c, err)
	}

	return c.JSON(http.StatusCreated, feature)
}

func (fc *featureController) DeleteFeature(c Context, id string) error {
	err := fc.featureInteractor.Delete(id)

	if err != nil {
		return sendErrorCar(c, err)
	}

	return c.JSON(http.StatusNoContent, nil)
}

func (fc *featureController) UpdateFeature(c Context) error {
	var params model.Feature

	if err := c.Bind(&params); err != nil {
		return sendErrorCar(c, model.NewAppError("Bad Request. "+err.Error(), http.StatusBadRequest))
	}

	feature, err := fc.featureInteractor.Update(&params)
	if err != nil {
		return sendErrorCar(c, err)
	}

	return c.JSON(http.StatusCreated, feature)
}
