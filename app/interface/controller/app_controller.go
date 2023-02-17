package controller

type AppController struct {
	Car interface{ CarController }
}
