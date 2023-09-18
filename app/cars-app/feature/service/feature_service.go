package service

import (
	dbr "api-cars/app/app-common/db"
	"api-cars/app/app-common/validator"
	"api-cars/app/cars-app/feature/model"
	"api-cars/app/cars-app/feature/presenter"
	"api-cars/app/cars-app/feature/repository"
	carError "api-cars/app/domain/model"

	"net/http"
	"strconv"
)

type featureService struct {
	FeatureRepository repository.FeatureRepository
	FeaturePresenter  presenter.FeaturePresenter
	DBRepository      dbr.DBRepository
}

type FeatureService interface {
	Get(f []*model.Feature) ([]*model.Feature, error)
	GetOne(id string) (*model.Feature, error)
	Create(f *model.Feature) (*model.Feature, error)
	Delete(id string) error
	Update(f *model.Feature) (*model.Feature, error)
}

func NewFeatureService(r repository.FeatureRepository, p presenter.FeaturePresenter, d dbr.DBRepository) FeatureService {
	return &featureService{r, p, d}
}

func (fi *featureService) Get(feature []*model.Feature) ([]*model.Feature, error) {
	feature, err := fi.FeatureRepository.FindAll(feature)
	if err != nil {
		return nil, err
	}

	return fi.FeaturePresenter.ResponseFeatures(feature), nil
}

func (fi *featureService) GetOne(id string) (*model.Feature, error) {

	idn, errValid := strconv.Atoi(id)
	if errValid != nil {
		return nil, errValid
	}

	feature, err := fi.FeatureRepository.FindOne(idn)

	if err != nil {
		return nil, carError.HandleError(err, "Feature with ID "+id+" not found. "+err.Error(), http.StatusNotFound)
	}

	return fi.FeaturePresenter.ResponseFeature(feature), nil
}

func (f *featureService) Create(feature *model.Feature) (*model.Feature, error) {

	err := validator.ValidateStruct(feature)

	if err != nil {
		return nil, carError.HandleError(err, err.Error(), http.StatusBadRequest)
	}

	data, err := f.DBRepository.Transaction(func(i interface{}) (interface{}, error) {
		feature, err := f.FeatureRepository.Create(feature)

		return feature, err
	})

	feature, ok := data.(*model.Feature)

	if !ok || err != nil {
		return nil, carError.HandleError(err, "Failed to create Feature.  "+err.Error(), http.StatusNotFound)
	}

	return feature, nil
}

func (fi *featureService) Delete(id string) error {

	err := fi.ValidateRecordExists(id)

	if err != nil {
		return carError.HandleError(err, "Feature with ID "+id+" not found. "+err.Error(), http.StatusNotFound)
	}

	idn, errValid := strconv.Atoi(id)
	if errValid != nil {
		return errValid
	}

	return fi.FeatureRepository.Delete(idn)
}

func (fi *featureService) Update(feature *model.Feature) (*model.Feature, error) {
	err := validator.ValidateStruct(feature)

	if err != nil {
		return nil, carError.HandleError(err, err.Error(), http.StatusBadRequest)
	}

	errExists := fi.ValidateRecordExists(strconv.Itoa((feature.Id)))

	if errExists != nil {
		return nil, carError.HandleError(errExists, "Feature with ID "+strconv.Itoa((feature.Id))+" not found. "+errExists.Error(), http.StatusNotFound)
	}

	data, err := fi.DBRepository.Transaction(func(i interface{}) (interface{}, error) {
		feature, err := fi.FeatureRepository.Update(feature)

		return feature, err
	})
	feature, ok := data.(*model.Feature)

	if !ok || err != nil {
		return nil, carError.HandleError(err, "Failed to update Feature.  "+err.Error(), http.StatusNotFound)
	}

	return feature, nil
}

func (fi *featureService) ValidateRecordExists(id string) error {
	feature, err := fi.GetOne(id)

	if err != nil && feature == nil {
		return err
	}

	return err
}
