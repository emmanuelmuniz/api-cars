package interactor

import (
	"api-cars/app/domain/model"
	"api-cars/app/usecase/presenter"
	"api-cars/app/usecase/repository"
	"net/http"
	"strconv"
)

type carModelInteractor struct {
	CarModelRepository repository.CarModelRepository
	MakeRepository     repository.MakeRepository
	CarModelPresenter  presenter.CarModelPresenter
	DBRepository       repository.DBRepository
}

type CarModelInteractor interface {
	Get(m []*model.CarModel) ([]*model.CarModel, error)
	GetOne(id string) (*model.CarModel, error)
	Create(m *model.CarModel) (*model.CarModel, error)
	Delete(id string) error
	Update(m *model.CarModel) (*model.CarModel, error)
}

func NewCarModelInteractor(r repository.CarModelRepository,
	mr repository.MakeRepository,
	p presenter.CarModelPresenter,
	d repository.DBRepository) CarModelInteractor {
	return &carModelInteractor{r, mr, p, d}
}

func (cmi *carModelInteractor) Get(carModel []*model.CarModel) ([]*model.CarModel, error) {
	carModel, err := cmi.CarModelRepository.FindAll(carModel)

	if err != nil {
		return nil, model.HandleError(err, "Error to retrieve records. "+err.Error(), http.StatusNotFound)
	}

	return cmi.CarModelPresenter.ResponseCarModels(carModel), nil
}

func (cmi *carModelInteractor) GetOne(id string) (*model.CarModel, error) {
	idn, errValid := strconv.Atoi(id)

	if errValid != nil {
		return nil, model.HandleError(errValid, "Bad request. "+errValid.Error(), http.StatusNotFound)
	}

	carModel, err := cmi.CarModelRepository.FindOne(idn)

	if err != nil {
		return nil, model.HandleError(err, "Car Model with ID "+id+" not found. "+err.Error(), http.StatusNotFound)
	}

	return cmi.CarModelPresenter.ResponseCarModel(carModel), nil
}

func (m *carModelInteractor) Create(carModel *model.CarModel) (*model.CarModel, error) {
	var err error
	var make *model.Make

	make, err = repository.MakeRepository.FindOne(m.MakeRepository, carModel.Make.Id)

	if err != nil {
		return nil, model.HandleError(err, "Make with ID "+strconv.Itoa(carModel.Make.Id)+" not found. "+err.Error(), http.StatusNotFound)
	}

	carModel.Make = *make

	data, err := m.DBRepository.Transaction(func(i interface{}) (interface{}, error) {
		carModel, err := m.CarModelRepository.Create(carModel)

		return carModel, err
	})

	carModel, ok := data.(*model.CarModel)

	if !ok {
		return nil, model.HandleError(err, "Failed to create Car Model. "+err.Error(), http.StatusBadRequest)
	}

	return carModel, nil
}

func (cmi *carModelInteractor) Delete(id string) error {

	idn, errValid := strconv.Atoi(id)

	if errValid != nil {
		return model.HandleError(errValid, "Bad request. "+errValid.Error(), http.StatusNotFound)
	}

	err := cmi.ValidateRecordExists(id)

	if err != nil {
		return model.HandleError(err, "Car Model with ID "+id+" not found. "+err.Error(), http.StatusNotFound)
	}

	return cmi.CarModelRepository.Delete(idn)
}

func (cmi *carModelInteractor) Update(carModel *model.CarModel) (*model.CarModel, error) {
	errExists := cmi.ValidateRecordExists(strconv.Itoa(carModel.Id))

	if errExists != nil {
		return nil, model.HandleError(errExists, "Car Model with ID "+strconv.Itoa(carModel.Id)+" not found. "+errExists.Error(), http.StatusNotFound)
	}

	var err error
	var make *model.Make

	make, err = repository.MakeRepository.FindOne(cmi.MakeRepository, carModel.Make.Id)

	if err != nil {
		return nil, model.HandleError(err, "Make with ID "+strconv.Itoa(carModel.Make.Id)+" not found. "+err.Error(), http.StatusNotFound)
	}

	carModel.Make = *make

	data, err := cmi.DBRepository.Transaction(func(i interface{}) (interface{}, error) {
		carModel, err := cmi.CarModelRepository.Update(carModel)

		return carModel, err
	})
	carModel, ok := data.(*model.CarModel)

	if !ok || err != nil {
		return nil, model.HandleError(err, "Failed to update Car Model.  "+err.Error(), http.StatusNotFound)
	}

	return carModel, nil
}

func (cmi *carModelInteractor) ValidateRecordExists(id string) error {
	carModel, err := cmi.GetOne(id)

	if err != nil && carModel == nil {
		return err
	}

	return err
}
