package interactor

import (
	"api-cars/app/domain/model"
	"api-cars/app/usecase/presenter"
	"api-cars/app/usecase/repository"
	"errors"
	"strconv"
)

type makeInteractor struct {
	MakeRepository repository.MakeRepository
	MakePresenter  presenter.MakePresenter
	DBRepository   repository.DBRepository
}

type MakeInteractor interface {
	Get(m []*model.Make) ([]*model.Make, error)
	GetOne(id string) (*model.Make, error)
	Create(m *model.Make) (*model.Make, error)
	Delete(id string) error
	Update(m *model.Make) (*model.Make, error)
}

func NewMakeInteractor(r repository.MakeRepository, p presenter.MakePresenter, d repository.DBRepository) MakeInteractor {
	return &makeInteractor{r, p, d}
}

func (mi *makeInteractor) Get(make []*model.Make) ([]*model.Make, error) {
	make, err := mi.MakeRepository.FindAll(make)
	if err != nil {
		return nil, err
	}

	return mi.MakePresenter.ResponseMakes(make), nil
}

func (mi *makeInteractor) GetOne(id string) (*model.Make, error) {

	idn, errValid := strconv.Atoi(id)
	if errValid != nil {
		return nil, errValid
	}

	make, err := mi.MakeRepository.FindOne(idn)

	if err != nil {
		return nil, err
	}

	return mi.MakePresenter.ResponseMake(make), nil
}

func (m *makeInteractor) Create(make *model.Make) (*model.Make, error) {
	data, err := m.DBRepository.Transaction(func(i interface{}) (interface{}, error) {
		make, err := m.MakeRepository.Create(make)

		return make, err
	})

	make, ok := data.(*model.Make)

	if !ok {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	return make, nil
}

func (mi *makeInteractor) Delete(id string) error {
	err := mi.ValidateRecordExists(id)

	idn, errValid := strconv.Atoi(id)
	if errValid != nil {
		return errValid
	}

	if err != nil {
		return err
	}

	return mi.MakeRepository.Delete(idn)
}

func (mi *makeInteractor) Update(make *model.Make) (*model.Make, error) {
	errExists := mi.ValidateRecordExists(strconv.Itoa((make.Id)))

	if errExists != nil {
		return nil, errExists
	}

	data, err := mi.DBRepository.Transaction(func(i interface{}) (interface{}, error) {
		make, err := mi.MakeRepository.Update(make)

		return make, err
	})
	make, ok := data.(*model.Make)

	if !ok {
		return nil, errors.New("Update error")
	}

	if err != nil {
		return nil, err
	}

	return make, nil
}

func (mi *makeInteractor) ValidateRecordExists(id string) error {
	make, err := mi.GetOne(id)

	if err != nil && make == nil {
		return err
	}

	return err
}
