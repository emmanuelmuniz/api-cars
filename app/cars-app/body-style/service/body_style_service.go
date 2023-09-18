package service

import (
	dbr "api-cars/app/app-common/db"
	carError "api-cars/app/domain/model"

	"api-cars/app/app-common/validator"
	"api-cars/app/cars-app/body-style/model"
	"api-cars/app/cars-app/body-style/presenter"
	"api-cars/app/cars-app/body-style/repository"

	"net/http"
	"strconv"
)

type bodyStyleService struct {
	BodyStyleRepository repository.BodyStyleRepository
	BodyStylePresenter  presenter.BodyStylePresenter
	DBRepository        dbr.DBRepository
}

type BodyStyleService interface {
	Get(bs []*model.BodyStyle) ([]*model.BodyStyle, error)
	GetOne(id string) (*model.BodyStyle, error)
	Create(bs *model.BodyStyle) (*model.BodyStyle, error)
	Delete(id string) error
	Update(bs *model.BodyStyle) (*model.BodyStyle, error)
}

func NewBodyStyleService(r repository.BodyStyleRepository, p presenter.BodyStylePresenter, d dbr.DBRepository) BodyStyleService {
	return &bodyStyleService{r, p, d}
}

func (bsi *bodyStyleService) Get(bodyStyle []*model.BodyStyle) ([]*model.BodyStyle, error) {
	bodyStyle, err := bsi.BodyStyleRepository.FindAll(bodyStyle)
	if err != nil {
		return nil, err
	}

	return bsi.BodyStylePresenter.ResponseBodyStyles(bodyStyle), nil
}

func (bsi *bodyStyleService) GetOne(id string) (*model.BodyStyle, error) {

	idn, errValid := strconv.Atoi(id)
	if errValid != nil {
		return nil, errValid
	}

	bodyStyle, err := bsi.BodyStyleRepository.FindOne(idn)

	if err != nil {
		return nil, carError.HandleError(err, "BodyStyle with ID "+id+" not found. "+err.Error(), http.StatusNotFound)
	}

	return bsi.BodyStylePresenter.ResponseBodyStyle(bodyStyle), nil
}

func (bs *bodyStyleService) Create(bodyStyle *model.BodyStyle) (*model.BodyStyle, error) {

	err := validator.ValidateStruct(bodyStyle)

	if err != nil {
		return nil, carError.HandleError(err, err.Error(), http.StatusBadRequest)
	}

	data, err := bs.DBRepository.Transaction(func(i interface{}) (interface{}, error) {
		bodyStyle, err := bs.BodyStyleRepository.Create(bodyStyle)

		return bodyStyle, err
	})

	bodyStyle, ok := data.(*model.BodyStyle)

	if !ok || err != nil {
		return nil, carError.HandleError(err, "Failed to create Body Style.  "+err.Error(), http.StatusNotFound)
	}

	return bodyStyle, nil
}

func (bsi *bodyStyleService) Delete(id string) error {
	err := bsi.ValidateRecordExists(id)

	if err != nil {
		return carError.HandleError(err, "BodyStyle with ID "+id+" not found. "+err.Error(), http.StatusNotFound)
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

func (bsi *bodyStyleService) Update(bodyStyle *model.BodyStyle) (*model.BodyStyle, error) {

	err := validator.ValidateStruct(bodyStyle)

	if err != nil {
		return nil, carError.HandleError(err, err.Error(), http.StatusBadRequest)
	}

	errExists := bsi.ValidateRecordExists(strconv.Itoa((bodyStyle.Id)))

	if errExists != nil {
		return nil, carError.HandleError(errExists, "BodyStyle with ID "+strconv.Itoa((bodyStyle.Id))+" not found. "+errExists.Error(), http.StatusNotFound)
	}

	data, err := bsi.DBRepository.Transaction(func(i interface{}) (interface{}, error) {
		bodyStyle, err := bsi.BodyStyleRepository.Update(bodyStyle)

		return bodyStyle, err
	})
	bodyStyle, ok := data.(*model.BodyStyle)

	if !ok || err != nil {
		return nil, carError.HandleError(err, "Failed to update Body Style.  "+err.Error(), http.StatusNotFound)
	}

	return bodyStyle, nil
}

func (bsi *bodyStyleService) ValidateRecordExists(id string) error {
	bodyStyle, err := bsi.GetOne(id)

	if err != nil && bodyStyle == nil {
		return err
	}

	return err
}
