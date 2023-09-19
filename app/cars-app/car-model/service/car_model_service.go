package service

import (
	dbr "api-cars/app/app-common/db"
	"api-cars/app/app-common/validator"
	"api-cars/app/cars-app/car-model/model"
	"api-cars/app/cars-app/car-model/presenter"
	"api-cars/app/cars-app/car-model/repository"
	mm "api-cars/app/cars-app/make/model"
	mr "api-cars/app/cars-app/make/repository"
	carError "api-cars/app/domain/model"

	"net/http"
	"strconv"
)

type carModelService struct {
	CarModelRepository repository.CarModelRepository
	MakeRepository     mr.MakeRepository
	CarModelPresenter  presenter.CarModelPresenter
	DBRepository       dbr.DBRepository
}

type CarModelService interface {
	Get(m []*model.CarModel) ([]*model.CarModel, error)
	GetOne(id string) (*model.CarModel, error)
	Create(m *model.CarModel) (*model.CarModel, error)
	Delete(id string) error
	Update(m *model.CarModel) (*model.CarModel, error)
}

func NewCarModelService(r repository.CarModelRepository,
	mr mr.MakeRepository,
	p presenter.CarModelPresenter,
	d dbr.DBRepository) CarModelService {
	return &carModelService{r, mr, p, d}
}

func (cmi *carModelService) Get(carModel []*model.CarModel) ([]*model.CarModel, error) {
	carModel, err := cmi.CarModelRepository.FindAll(carModel)

	if err != nil {
		return nil, carError.HandleError(err, "Error to retrieve records. "+err.Error(), http.StatusNotFound)
	}

	return cmi.CarModelPresenter.ResponseCarModels(carModel), nil
}

func (cmi *carModelService) GetOne(id string) (*model.CarModel, error) {
	idn, errValid := strconv.Atoi(id)

	if errValid != nil {
		return nil, carError.HandleError(errValid, "Bad request. "+errValid.Error(), http.StatusNotFound)
	}

	carModel, err := cmi.CarModelRepository.FindOne(idn)

	if err != nil {
		return nil, carError.HandleError(err, "Car Model with ID "+id+" not found. "+err.Error(), http.StatusNotFound)
	}

	return cmi.CarModelPresenter.ResponseCarModel(carModel), nil
}

func (m *carModelService) Create(carModel *model.CarModel) (*model.CarModel, error) {
	var err error
	var make *mm.Make

	err = validator.ValidateStruct(carModel)

	if err != nil {
		return nil, carError.HandleError(err, err.Error(), http.StatusBadRequest)
	}

	make, err = mr.MakeRepository.FindOne(m.MakeRepository, carModel.MakeID)

	if err != nil || make == nil {
		return nil, carError.HandleError(err, "Make with ID "+strconv.Itoa(carModel.MakeID)+" not found. "+err.Error(), http.StatusNotFound)
	}

	data, err := m.DBRepository.Transaction(func(i interface{}) (interface{}, error) {
		carModel, err := m.CarModelRepository.Create(carModel)

		return carModel, err
	})

	carModel, ok := data.(*model.CarModel)

	if !ok {
		return nil, carError.HandleError(err, "Failed to create Car Model. "+err.Error(), http.StatusBadRequest)
	}

	return carModel, nil
}

func (cmi *carModelService) Delete(id string) error {

	idn, errValid := strconv.Atoi(id)

	if errValid != nil {
		return carError.HandleError(errValid, "Bad request. "+errValid.Error(), http.StatusNotFound)
	}

	err := cmi.ValidateRecordExists(id)

	if err != nil {
		return carError.HandleError(err, "Car Model with ID "+id+" not found. "+err.Error(), http.StatusNotFound)
	}

	return cmi.CarModelRepository.Delete(idn)
}

func (cmi *carModelService) Update(carModel *model.CarModel) (*model.CarModel, error) {
	var err error

	err = validator.ValidateStruct(carModel)

	if err != nil {
		return nil, carError.HandleError(err, err.Error(), http.StatusBadRequest)
	}

	errExists := cmi.ValidateRecordExists(strconv.Itoa(carModel.Id))

	if errExists != nil {
		return nil, carError.HandleError(errExists, "Car Model with ID "+strconv.Itoa(carModel.Id)+" not found. "+errExists.Error(), http.StatusNotFound)
	}

	var make *mm.Make

	make, err = mr.MakeRepository.FindOne(cmi.MakeRepository, carModel.MakeID)

	if err != nil || make == nil {
		return nil, carError.HandleError(err, "Make with ID "+strconv.Itoa(carModel.MakeID)+" not found. "+err.Error(), http.StatusNotFound)
	}

	data, err := cmi.DBRepository.Transaction(func(i interface{}) (interface{}, error) {
		carModel, err := cmi.CarModelRepository.Update(carModel)

		return carModel, err
	})
	carModel, ok := data.(*model.CarModel)

	if !ok || err != nil {
		return nil, carError.HandleError(err, "Failed to update Car Model.  "+err.Error(), http.StatusNotFound)
	}

	return carModel, nil
}

func (cmi *carModelService) ValidateRecordExists(id string) error {
	carModel, err := cmi.GetOne(id)

	if err != nil && carModel == nil {
		return err
	}

	return err
}
