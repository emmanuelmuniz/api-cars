package interactor

import (
	"api-cars/app/domain/model"
	"api-cars/app/usecase/presenter"
	"api-cars/app/usecase/repository"
	"errors"
)

type carModelInteractor struct {
	CarModelRepository repository.CarModelRepository
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

func NewCarModelInteractor(r repository.CarModelRepository, p presenter.CarModelPresenter, d repository.DBRepository) CarModelInteractor {
	return &carModelInteractor{r, p, d}
}

func (cmi *carModelInteractor) Get(carModel []*model.CarModel) ([]*model.CarModel, error) {
	carModel, err := cmi.CarModelRepository.FindAll(carModel)
	if err != nil {
		return nil, err
	}

	return cmi.CarModelPresenter.ResponseCarModels(carModel), nil
}

func (cmi *carModelInteractor) GetOne(id string) (*model.CarModel, error) {
	carModel, err := cmi.CarModelRepository.FindOne(id)

	if err != nil {
		return nil, err
	}

	return cmi.CarModelPresenter.ResponseCarModel(carModel), nil
}

func (m *carModelInteractor) Create(carModel *model.CarModel) (*model.CarModel, error) {
	data, err := m.DBRepository.Transaction(func(i interface{}) (interface{}, error) {
		carModel, err := m.CarModelRepository.Create(carModel)

		return carModel, err
	})

	carModel, ok := data.(*model.CarModel)

	if !ok {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	return carModel, nil
}

func (cmi *carModelInteractor) Delete(id string) error {
	err := cmi.ValidateRecordExists(id)

	if err != nil {
		return err
	}

	return cmi.CarModelRepository.Delete(id)
}

func (cmi *carModelInteractor) Update(carModel *model.CarModel) (*model.CarModel, error) {
	errExists := cmi.ValidateRecordExists(string(carModel.Id))

	if errExists != nil {
		return nil, errExists
	}

	data, err := cmi.DBRepository.Transaction(func(i interface{}) (interface{}, error) {
		carModel, err := cmi.CarModelRepository.Update(carModel)

		return carModel, err
	})
	carModel, ok := data.(*model.CarModel)

	if !ok {
		return nil, errors.New("Update error")
	}

	if err != nil {
		return nil, err
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
