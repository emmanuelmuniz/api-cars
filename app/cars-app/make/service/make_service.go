package service

import (
	dbr "api-cars/app/app-common/db"
	"api-cars/app/app-common/validator"
	"api-cars/app/cars-app/make/model"
	"api-cars/app/cars-app/make/presenter"
	"api-cars/app/cars-app/make/repository"
	carError "api-cars/app/domain/model"

	"net/http"
	"strconv"
)

type makeService struct {
	MakeRepository repository.MakeRepository
	MakePresenter  presenter.MakePresenter
	DBRepository   dbr.DBRepository
}

type MakeService interface {
	Get(m []*model.Make) ([]*model.Make, error)
	GetOne(id string) (*model.Make, error)
	Create(m *model.Make) (*model.Make, error)
	Delete(id string) error
	Update(m *model.Make) (*model.Make, error)
}

func NewMakeService(r repository.MakeRepository, p presenter.MakePresenter, d dbr.DBRepository) MakeService {
	return &makeService{r, p, d}
}

func (mi *makeService) Get(make []*model.Make) ([]*model.Make, error) {
	make, err := mi.MakeRepository.FindAll(make)
	if err != nil {
		return nil, err
	}

	return mi.MakePresenter.ResponseMakes(make), nil
}

func (mi *makeService) GetOne(id string) (*model.Make, error) {

	idn, errValid := strconv.Atoi(id)
	if errValid != nil {
		return nil, errValid
	}

	make, err := mi.MakeRepository.FindOne(idn)

	if err != nil {
		return nil, carError.HandleError(err, "Make with ID "+id+" not found. "+err.Error(), http.StatusNotFound)
	}

	return mi.MakePresenter.ResponseMake(make), nil
}

func (m *makeService) Create(make *model.Make) (*model.Make, error) {

	err := validator.ValidateStruct(make)

	if err != nil {
		return nil, carError.HandleError(err, err.Error(), http.StatusBadRequest)
	}

	data, err := m.DBRepository.Transaction(func(i interface{}) (interface{}, error) {
		make, err := m.MakeRepository.Create(make)

		return make, err
	})

	make, ok := data.(*model.Make)

	if !ok || err != nil {
		return nil, carError.HandleError(err, "Failed to create Make.  "+err.Error(), http.StatusNotFound)
	}

	return make, nil
}

func (mi *makeService) Delete(id string) error {
	err := mi.ValidateRecordExists(id)

	if err != nil {
		return carError.HandleError(err, "Make with ID "+id+" not found. "+err.Error(), http.StatusNotFound)
	}

	idn, errValid := strconv.Atoi(id)
	if errValid != nil {
		return errValid
	}

	if err != nil {
		return err
	}

	return mi.MakeRepository.Delete(idn)
}

func (mi *makeService) Update(make *model.Make) (*model.Make, error) {
	err := validator.ValidateStruct(make)

	if err != nil {
		return nil, carError.HandleError(err, err.Error(), http.StatusBadRequest)
	}

	errExists := mi.ValidateRecordExists(strconv.Itoa((make.Id)))

	if errExists != nil {
		return nil, carError.HandleError(errExists, "Make with ID "+strconv.Itoa((make.Id))+" not found. "+errExists.Error(), http.StatusNotFound)
	}

	data, err := mi.DBRepository.Transaction(func(i interface{}) (interface{}, error) {
		make, err := mi.MakeRepository.Update(make)

		return make, err
	})
	make, ok := data.(*model.Make)

	if !ok || err != nil {
		return nil, carError.HandleError(err, "Failed to update Make.  "+err.Error(), http.StatusNotFound)
	}

	return make, nil
}

func (mi *makeService) ValidateRecordExists(id string) error {
	make, err := mi.GetOne(id)

	if err != nil && make == nil {
		return err
	}

	return err
}
