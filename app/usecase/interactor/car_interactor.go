package interactor

import (
	"api-cars/app/domain/model"
	"api-cars/app/usecase/presenter"
	"api-cars/app/usecase/repository"
	"net/http"
	"strconv"
)

type carInteractor struct {
	CarRepository       repository.CarRepository
	CarPresenter        presenter.CarPresenter
	MakeRepository      repository.MakeRepository
	CarModelRepository  repository.CarModelRepository
	BodyStyleRepository repository.BodyStyleRepository
	DBRepository        repository.DBRepository
}

type CarInteractor interface {
	Get(c []*model.Car) ([]*model.Car, error)
	GetOne(id string) (*model.Car, error)
	Create(c *model.Car) (*model.Car, error)
	Delete(id string) error
	Update(c *model.Car) (*model.Car, error)
}

func NewCarInteractor(r repository.CarRepository,
	p presenter.CarPresenter,
	mr repository.MakeRepository,
	cmr repository.CarModelRepository,
	bdr repository.BodyStyleRepository,
	d repository.DBRepository) CarInteractor {
	return &carInteractor{r, p, mr, cmr, bdr, d}
}

func (ci *carInteractor) Get(car []*model.Car) ([]*model.Car, error) {
	car, err := ci.CarRepository.FindAll(car)

	if err != nil {
		return nil, model.HandleError(err, "Error to retrieve records. "+err.Error(), http.StatusNotFound)
	}

	return ci.CarPresenter.ResponseCars(car), nil
}

func (ci *carInteractor) GetOne(id string) (*model.Car, error) {

	idn, errValid := strconv.Atoi(id)
	if errValid != nil {
		return nil, errValid
	}

	car, err := ci.CarRepository.FindOne(idn)

	if err != nil {
		return nil, model.HandleError(err, "Car with ID "+id+" not found. "+err.Error(), http.StatusNotFound)
	}

	return ci.CarPresenter.ResponseCar(car), nil
}

func (c *carInteractor) Create(car *model.Car) (*model.Car, error) {
	var err error
	var carModel *model.CarModel

	carModel, err = repository.CarModelRepository.FindOne(c.CarModelRepository, car.CarModel.Id)

	if err != nil {
		return nil, model.HandleError(err, "Car Model with ID "+strconv.Itoa(car.CarModel.Id)+" not found. "+err.Error(), http.StatusNotFound)
	}

	car.Make = carModel.Make
	var newCarmodel = carModel
	car.CarModel = *newCarmodel

	var bodyStyle *model.BodyStyle

	bodyStyle, err = repository.BodyStyleRepository.FindOne(c.BodyStyleRepository, car.BodyStyle.Id)

	if err != nil {
		return nil, model.HandleError(err, "Body Style with ID "+strconv.Itoa(car.CarModel.Id)+" not found. "+err.Error(), http.StatusNotFound)
	}

	car.BodyStyle = *bodyStyle

	data, err := c.DBRepository.Transaction(func(i interface{}) (interface{}, error) {
		car, err := c.CarRepository.Create(car)

		return car, err
	})

	car, ok := data.(*model.Car)

	if !ok || err != nil {
		return nil, model.HandleError(err, "Failed to create Car. "+err.Error(), http.StatusNotModified)
	}

	return car, nil
}

func (ci *carInteractor) Delete(id string) error {
	idn, errValid := strconv.Atoi(id)

	if errValid != nil {
		return model.HandleError(errValid, "Bad Request. "+errValid.Error(), http.StatusBadRequest)
	}

	errExists := ci.ValidateRecordExists(id)

	if errExists != nil {
		return model.HandleError(errExists, "Car with ID "+id+" not found. "+errExists.Error(), http.StatusNotFound)
	}

	return ci.CarRepository.Delete(idn)
}

func (ci *carInteractor) Update(car *model.Car) (*model.Car, error) {
	errExists := ci.ValidateRecordExists(strconv.Itoa(car.Id))

	if errExists != nil {
		return nil, model.HandleError(errExists, "Car with ID "+strconv.Itoa(car.Id)+" not found. "+errExists.Error(), http.StatusNotFound)
	}

	var err error
	var carModel *model.CarModel

	carModel, err = repository.CarModelRepository.FindOne(ci.CarModelRepository, car.CarModel.Id)

	if err != nil {
		return nil, model.HandleError(errExists, "Car Model with ID "+strconv.Itoa(car.CarModel.Id)+" not found. "+err.Error(), http.StatusNotFound)
	}

	var bodyStyle *model.BodyStyle

	bodyStyle, err = repository.BodyStyleRepository.FindOne(ci.BodyStyleRepository, car.BodyStyle.Id)

	if err != nil {
		return nil, model.HandleError(errExists, "Body Style with ID "+strconv.Itoa(car.CarModel.Id)+" not found. "+err.Error(), http.StatusNotFound)
	}

	car.Make = carModel.Make
	var newCarmodel = carModel
	car.CarModel = *newCarmodel

	car.BodyStyle = *bodyStyle

	data, err := ci.DBRepository.Transaction(func(i interface{}) (interface{}, error) {
		car, err := ci.CarRepository.Update(car)

		return car, err
	})
	car, ok := data.(*model.Car)

	if !ok || err != nil {
		return nil, model.HandleError(errExists, "Failed to update Car. "+err.Error(), http.StatusNotModified)
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
