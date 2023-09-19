package service

import (
	dbr "api-cars/app/app-common/db"
	"api-cars/app/app-common/validator"
	"api-cars/app/cars-app/car-image/model"
	"api-cars/app/cars-app/car-image/presenter"
	"api-cars/app/cars-app/car-image/repository"
	carError "api-cars/app/domain/model"

	"net/http"
	"strconv"
)

type carImageService struct {
	CarImageRepository repository.CarImageRepository
	CarImagePresenter  presenter.CarImagePresenter
	DBRepository       dbr.DBRepository
}

type CarImageService interface {
	Get(f []*model.CarImage) ([]*model.CarImage, error)
	GetOne(id string) (*model.CarImage, error)
	Create(f *model.CarImage) (*model.CarImage, error)
	Delete(id string) error
	Update(f *model.CarImage) (*model.CarImage, error)
}

func NewCarImageService(r repository.CarImageRepository, p presenter.CarImagePresenter, d dbr.DBRepository) CarImageService {
	return &carImageService{r, p, d}
}

func (fi *carImageService) Get(carImage []*model.CarImage) ([]*model.CarImage, error) {
	carImage, err := fi.CarImageRepository.FindAll(carImage)
	if err != nil {
		return nil, err
	}

	return fi.CarImagePresenter.ResponseCarImages(carImage), nil
}

func (fi *carImageService) GetOne(id string) (*model.CarImage, error) {

	idn, errValid := strconv.Atoi(id)
	if errValid != nil {
		return nil, errValid
	}

	carImage, err := fi.CarImageRepository.FindOne(idn)

	if err != nil {
		return nil, carError.HandleError(err, "CarImage with ID "+id+" not found. "+err.Error(), http.StatusNotFound)
	}

	return fi.CarImagePresenter.ResponseCarImage(carImage), nil
}

func (f *carImageService) Create(carImage *model.CarImage) (*model.CarImage, error) {

	err := validator.ValidateStruct(carImage)

	if err != nil {
		return nil, carError.HandleError(err, err.Error(), http.StatusBadRequest)
	}

	data, err := f.DBRepository.Transaction(func(i interface{}) (interface{}, error) {
		carImage, err := f.CarImageRepository.Create(carImage)

		return carImage, err
	})

	carImage, ok := data.(*model.CarImage)

	if !ok || err != nil {
		return nil, carError.HandleError(err, "Failed to create CarImage.  "+err.Error(), http.StatusNotFound)
	}

	return carImage, nil
}

func (fi *carImageService) Delete(id string) error {

	err := fi.ValidateRecordExists(id)

	if err != nil {
		return carError.HandleError(err, "CarImage with ID "+id+" not found. "+err.Error(), http.StatusNotFound)
	}

	idn, errValid := strconv.Atoi(id)
	if errValid != nil {
		return errValid
	}

	return fi.CarImageRepository.Delete(idn)
}

func (fi *carImageService) Update(carImage *model.CarImage) (*model.CarImage, error) {
	err := validator.ValidateStruct(carImage)

	if err != nil {
		return nil, carError.HandleError(err, err.Error(), http.StatusBadRequest)
	}

	errExists := fi.ValidateRecordExists(strconv.Itoa((carImage.Id)))

	if errExists != nil {
		return nil, carError.HandleError(errExists, "CarImage with ID "+strconv.Itoa((carImage.Id))+" not found. "+errExists.Error(), http.StatusNotFound)
	}

	data, err := fi.DBRepository.Transaction(func(i interface{}) (interface{}, error) {
		carImage, err := fi.CarImageRepository.Update(carImage)

		return carImage, err
	})
	carImage, ok := data.(*model.CarImage)

	if !ok || err != nil {
		return nil, carError.HandleError(err, "Failed to update CarImage.  "+err.Error(), http.StatusNotFound)
	}

	return carImage, nil
}

func (fi *carImageService) ValidateRecordExists(id string) error {
	carImage, err := fi.GetOne(id)

	if err != nil && carImage == nil {
		return err
	}

	return err
}
