package interactor

import (
	"api-cars/app/domain/model"
	"api-cars/app/usecase/presenter"
	"api-cars/app/usecase/repository"
	"net/http"
	"strconv"
)

type featureInteractor struct {
	FeatureRepository repository.FeatureRepository
	FeaturePresenter  presenter.FeaturePresenter
	DBRepository      repository.DBRepository
}

type FeatureInteractor interface {
	Get(f []*model.Feature) ([]*model.Feature, error)
	GetOne(id string) (*model.Feature, error)
	Create(f *model.Feature) (*model.Feature, error)
	Delete(id string) error
	Update(f *model.Feature) (*model.Feature, error)
}

func NewFeatureInteractor(r repository.FeatureRepository, p presenter.FeaturePresenter, d repository.DBRepository) FeatureInteractor {
	return &featureInteractor{r, p, d}
}

func (fi *featureInteractor) Get(feature []*model.Feature) ([]*model.Feature, error) {
	feature, err := fi.FeatureRepository.FindAll(feature)
	if err != nil {
		return nil, err
	}

	return fi.FeaturePresenter.ResponseFeatures(feature), nil
}

func (fi *featureInteractor) GetOne(id string) (*model.Feature, error) {

	idn, errValid := strconv.Atoi(id)
	if errValid != nil {
		return nil, errValid
	}

	feature, err := fi.FeatureRepository.FindOne(idn)

	if err != nil {
		return nil, model.HandleError(err, "Feature with ID "+id+" not found. "+err.Error(), http.StatusNotFound)
	}

	return fi.FeaturePresenter.ResponseFeature(feature), nil
}

func (f *featureInteractor) Create(feature *model.Feature) (*model.Feature, error) {
	data, err := f.DBRepository.Transaction(func(i interface{}) (interface{}, error) {
		feature, err := f.FeatureRepository.Create(feature)

		return feature, err
	})

	feature, ok := data.(*model.Feature)

	if !ok || err != nil {
		return nil, model.HandleError(err, "Failed to create Feature.  "+err.Error(), http.StatusNotFound)
	}

	return feature, nil
}

func (fi *featureInteractor) Delete(id string) error {

	err := fi.ValidateRecordExists(id)

	if err != nil {
		return model.HandleError(err, "Feature with ID "+id+" not found. "+err.Error(), http.StatusNotFound)
	}

	idn, errValid := strconv.Atoi(id)
	if errValid != nil {
		return errValid
	}

	return fi.FeatureRepository.Delete(idn)
}

func (fi *featureInteractor) Update(feature *model.Feature) (*model.Feature, error) {
	errExists := fi.ValidateRecordExists(strconv.Itoa((feature.Id)))

	if errExists != nil {
		return nil, model.HandleError(errExists, "Feature with ID "+strconv.Itoa((feature.Id))+" not found. "+errExists.Error(), http.StatusNotFound)
	}

	data, err := fi.DBRepository.Transaction(func(i interface{}) (interface{}, error) {
		feature, err := fi.FeatureRepository.Update(feature)

		return feature, err
	})
	feature, ok := data.(*model.Feature)

	if !ok || err != nil {
		return nil, model.HandleError(err, "Failed to update Feature.  "+err.Error(), http.StatusNotFound)
	}

	return feature, nil
}

func (fi *featureInteractor) ValidateRecordExists(id string) error {
	feature, err := fi.GetOne(id)

	if err != nil && feature == nil {
		return err
	}

	return err
}
