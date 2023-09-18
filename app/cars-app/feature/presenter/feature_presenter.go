package presenter

import (
	"api-cars/app/cars-app/feature/model"
)

type featurePresenter struct{}

func NewFeaturePresenter() FeaturePresenter {
	return &featurePresenter{}
}

func (fp *featurePresenter) ResponseFeatures(features []*model.Feature) []*model.Feature {
	return features
}

func (fp *featurePresenter) ResponseFeature(feature *model.Feature) *model.Feature {
	return feature
}
