package interactor

import (
	"api-cars/app/domain/model"
	"api-cars/app/usecase/presenter"
	"api-cars/app/usecase/repository"
	"net/http"
	"strconv"
)

type bodyStyleInteractor struct {
	BodyStyleRepository repository.BodyStyleRepository
	BodyStylePresenter  presenter.BodyStylePresenter
	DBRepository        repository.DBRepository
}

type BodyStyleInteractor interface {
	Get(bs []*model.BodyStyle) ([]*model.BodyStyle, error)
	GetOne(id string) (*model.BodyStyle, error)
	Create(bs *model.BodyStyle) (*model.BodyStyle, error)
	Delete(id string) error
	Update(bs *model.BodyStyle) (*model.BodyStyle, error)
}

func NewBodyStyleInteractor(r repository.BodyStyleRepository, p presenter.BodyStylePresenter, d repository.DBRepository) BodyStyleInteractor {
	return &bodyStyleInteractor{r, p, d}
}

func (bsi *bodyStyleInteractor) Get(bodyStyle []*model.BodyStyle) ([]*model.BodyStyle, error) {
	bodyStyle, err := bsi.BodyStyleRepository.FindAll(bodyStyle)
	if err != nil {
		return nil, err
	}

	return bsi.BodyStylePresenter.ResponseBodyStyles(bodyStyle), nil
}

func (bsi *bodyStyleInteractor) GetOne(id string) (*model.BodyStyle, error) {

	idn, errValid := strconv.Atoi(id)
	if errValid != nil {
		return nil, errValid
	}

	bodyStyle, err := bsi.BodyStyleRepository.FindOne(idn)

	if err != nil {
		return nil, err
	}

	return bsi.BodyStylePresenter.ResponseBodyStyle(bodyStyle), nil
}

func (bs *bodyStyleInteractor) Create(bodyStyle *model.BodyStyle) (*model.BodyStyle, error) {
	data, err := bs.DBRepository.Transaction(func(i interface{}) (interface{}, error) {
		bodyStyle, err := bs.BodyStyleRepository.Create(bodyStyle)

		return bodyStyle, err
	})

	bodyStyle, ok := data.(*model.BodyStyle)

	if !ok || err != nil {
		return nil, model.HandleError(err, "Failed to create Body Style.  "+err.Error(), http.StatusNotFound)
	}

	return bodyStyle, nil
}

func (bsi *bodyStyleInteractor) Delete(id string) error {
	err := bsi.ValidateRecordExists(id)

	if err != nil {
		return model.HandleError(err, "BodyStyle with ID "+id+" not found. "+err.Error(), http.StatusNotFound)
	}

	idn, errValid := strconv.Atoi(id)
	if errValid != nil {
		return errValid
	}

	if err != nil {
		return err
	}

	return bsi.BodyStyleRepository.Delete(idn)
}

func (bsi *bodyStyleInteractor) Update(bodyStyle *model.BodyStyle) (*model.BodyStyle, error) {
	errExists := bsi.ValidateRecordExists(strconv.Itoa((bodyStyle.Id)))

	if errExists != nil {
		return nil, model.HandleError(errExists, "BodyStyle with ID "+strconv.Itoa((bodyStyle.Id))+" not found. "+errExists.Error(), http.StatusNotFound)
	}

	data, err := bsi.DBRepository.Transaction(func(i interface{}) (interface{}, error) {
		bodyStyle, err := bsi.BodyStyleRepository.Update(bodyStyle)

		return bodyStyle, err
	})
	bodyStyle, ok := data.(*model.BodyStyle)

	if !ok || err != nil {
		return nil, model.HandleError(err, "Failed to update Body Style.  "+err.Error(), http.StatusNotFound)
	}

	return bodyStyle, nil
}

func (bsi *bodyStyleInteractor) ValidateRecordExists(id string) error {
	bodyStyle, err := bsi.GetOne(id)

	if err != nil && bodyStyle == nil {
		return err
	}

	return err
}
