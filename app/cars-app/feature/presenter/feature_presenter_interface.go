package presenter

import "api-cars/app/cars-app/feature/model"

type FeaturePresenter interface {
	ResponseFeatures(f []*model.Feature) []*model.Feature
	ResponseFeature(f *model.Feature) *model.Feature
}
