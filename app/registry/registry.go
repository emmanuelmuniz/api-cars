package registry

import (
	"api-cars/app/interface/controller"

	"gorm.io/gorm"
)

type registry struct {
	db *gorm.DB
}

type Registry interface {
	NewAppController() controller.AppController
}

func NewRegistry(db *gorm.DB) Registry {
	return &registry{db}
}

func (r *registry) NewAppController() controller.AppController {
	return controller.AppController{
		Car:       r.NewCarController(),
		Make:      r.NewMakeController(),
		CarModel:  r.NewCarModelController(),
		BodyStyle: r.NewBodyStyleController(),
		Feature:   r.NewFeatureController(),
	}
}
