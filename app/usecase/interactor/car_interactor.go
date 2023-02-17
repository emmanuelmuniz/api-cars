package interactor

import (
	"api-cars/app/domain/model"
	"api-cars/app/usecase/presenter"
	"api-cars/app/usecase/repository"
	"errors"
)

type carInteractor struct {
	CarRepository repository.CarRepository
	CarPresenter  presenter.CarPresenter
	DBRepository  repository.DBRepository
}

type CarInteractor interface {
	Get(u []*model.Car) ([]*model.Car, error)
	Create(u *model.Car) (*model.Car, error)
}

func NewCarInteractor(r repository.CarRepository, p presenter.CarPresenter, d repository.DBRepository) CarInteractor {
	return &carInteractor{r, p, d}
}

func (us *carInteractor) Get(u []*model.Car) ([]*model.Car, error) {
	u, err := us.CarRepository.FindAll(u)
	if err != nil {
		return nil, err
	}

	return us.CarPresenter.ResponseCars(u), nil
}

func (c *carInteractor) Create(car *model.Car) (*model.Car, error) {
	data, err := c.DBRepository.Transaction(func(i interface{}) (interface{}, error) {
		u, err := c.CarRepository.Create(car)

		// do mailing
		// do logging
		// do another process
		return u, err
	})
	car, ok := data.(*model.Car)

	if !ok {
		return nil, errors.New("cast error")
	}

	if err != nil {
		return nil, err
	}

	return car, nil
}
