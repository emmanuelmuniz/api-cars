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
	Get(c []*model.Car) ([]*model.Car, error)
	GetOne(id string) (*model.Car, error)
	Create(c *model.Car) (*model.Car, error)
	Delete(id string) error
	Update(c *model.Car) (*model.Car, error)
}

func NewCarInteractor(r repository.CarRepository, p presenter.CarPresenter, d repository.DBRepository) CarInteractor {
	return &carInteractor{r, p, d}
}

func (ci *carInteractor) Get(car []*model.Car) ([]*model.Car, error) {
	car, err := ci.CarRepository.FindAll(car)
	if err != nil {
		return nil, err
	}

	return ci.CarPresenter.ResponseCars(car), nil
}

func (ci *carInteractor) GetOne(id string) (*model.Car, error) {
	car, err := ci.CarRepository.FindOne(id)

	if err != nil {
		return nil, err
	}

	return ci.CarPresenter.ResponseCar(car), nil
}

func (c *carInteractor) Create(car *model.Car) (*model.Car, error) {
	data, err := c.DBRepository.Transaction(func(i interface{}) (interface{}, error) {
		car, err := c.CarRepository.Create(car)

		return car, err
	})

	car, ok := data.(*model.Car)

	if !ok {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	return car, nil
}

func (ci *carInteractor) Delete(id string) error {
	err := ci.ValidateRecordExists(id)

	if err != nil {
		return err
	}

	return ci.CarRepository.Delete(id)
}

func (ci *carInteractor) Update(car *model.Car) (*model.Car, error) {
	errExists := ci.ValidateRecordExists(string(car.Id))

	if errExists != nil {
		return nil, errExists
	}

	data, err := ci.DBRepository.Transaction(func(i interface{}) (interface{}, error) {
		car, err := ci.CarRepository.Update(car)

		return car, err
	})
	car, ok := data.(*model.Car)

	if !ok {
		return nil, errors.New("Update error")
	}

	if err != nil {
		return nil, err
	}

	return car, nil
}

func (ci *carInteractor) ValidateRecordExists(id string) error {
	car, err := ci.GetOne(id)

	if err != nil && car == nil {
		return err
	}

	return err
}
