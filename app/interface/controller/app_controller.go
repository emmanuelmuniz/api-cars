package controller

type AppController struct {
	Car       interface{ CarController }
	Make      interface{ MakeController }
	CarModel  interface{ CarModelController }
	BodyStyle interface{ BodyStyleController }
	Feature   interface{ FeatureController }
}
