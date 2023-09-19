package controller

import (
	bsc "api-cars/app/cars-app/body-style/controller"
	cic "api-cars/app/cars-app/car-image/controller"
	cmc "api-cars/app/cars-app/car-model/controller"
	cc "api-cars/app/cars-app/car/controller"
	fc "api-cars/app/cars-app/feature/controller"
	mc "api-cars/app/cars-app/make/controller"
)

type AppController struct {
	Car       interface{ cc.CarController }
	Make      interface{ mc.MakeController }
	CarModel  interface{ cmc.CarModelController }
	BodyStyle interface{ bsc.BodyStyleController }
	Feature   interface{ fc.FeatureController }
	CarImage  interface{ cic.CarImageController }
}
