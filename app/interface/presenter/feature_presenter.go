package presenter

import (
	"api-cars/app/domain/model"
	"api-cars/app/usecase/presenter"
)

type featurePresenter struct{}

func NewFeaturePresenter() presenter.FeaturePresenter {
	return &featurePresenter{}
}

func (fp *featurePresenter) ResponseFeatures(features []*model.Feature) []*model.Feature {
	return features
}

func (fp *featurePresenter) ResponseFeature(feature *model.Feature) *model.Feature {
	return feature
}
