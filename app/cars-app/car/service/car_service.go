package service

import (
	"api-cars/app/cars-app/car/repository"

	dbr "api-cars/app/app-common/db"
	bsr "api-cars/app/cars-app/body-style/repository"
	cmr "api-cars/app/cars-app/car-model/repository"
	fr "api-cars/app/cars-app/feature/repository"
	mr "api-cars/app/cars-app/make/repository"

	"api-cars/app/cars-app/car/model"

	bsm "api-cars/app/cars-app/body-style/model"
	cmm "api-cars/app/cars-app/car-model/model"
	fm "api-cars/app/cars-app/feature/model"
	mm "api-cars/app/cars-app/make/model"
	carError "api-cars/app/domain/model"

	"api-cars/app/app-common/validator"
	"api-cars/app/cars-app/car/presenter"

	"net/http"
	"strconv"
)

type carService struct {
	CarRepository       repository.CarRepository
	CarPresenter        presenter.CarPresenter
	MakeRepository      mr.MakeRepository
	CarModelRepository  cmr.CarModelRepository
	BodyStyleRepository bsr.BodyStyleRepository
	FeatureRepository   fr.FeatureRepository
	DBRepository        dbr.DBRepository
}

type CarService interface {
	Get(c []*model.Car) ([]*model.Car, error)
	GetOne(id string) (*model.Car, error)
	Create(c *model.Car) (*model.Car, error)
	Delete(id string) error
	Update(c *model.Car) (*model.Car, error)
}

func NewCarService(r repository.CarRepository,
	carPresenter presenter.CarPresenter,
	makeRepository mr.MakeRepository,
	carModelRepository cmr.CarModelRepository,
	bodyStyleRepository bsr.BodyStyleRepository,
	featureRepository fr.FeatureRepository,
	databaseRepository dbr.DBRepository) CarService {
	return &carService{r, carPresenter, makeRepository, carModelRepository, bodyStyleRepository, featureRepository, databaseRepository}
}

func (ci *carService) Get(car []*model.Car) ([]*model.Car, error) {
	car, err := ci.CarRepository.FindAll(car)

	if err != nil {
		return nil, carError.HandleError(err, "Error to retrieve records. "+err.Error(), http.StatusNotFound)
	}

	return ci.CarPresenter.ResponseCars(car), nil
}

func (ci *carService) GetOne(id string) (*model.Car, error) {

	idn, errValid := strconv.Atoi(id)
	if errValid != nil {
		return nil, errValid
	}

	car, err := ci.CarRepository.FindOne(idn)

	if err != nil {
		return nil, carError.HandleError(err, "Car with ID "+id+" not found. "+err.Error(), http.StatusNotFound)
	}

	return ci.CarPresenter.ResponseCar(car), nil
}

func (c *carService) Create(car *model.Car) (*model.Car, error) {
	var err error

	err = validator.ValidateStruct(car)

	if err != nil {
		return nil, carError.HandleError(err, err.Error(), http.StatusBadRequest)
	}

	var features []*fm.Feature
	var featuresIDs []int

	for _, d := range car.Features {
		featuresIDs = append(featuresIDs, d.Id)
	}

	features, err = fr.FeatureRepository.FindByIDs(c.FeatureRepository, featuresIDs)

	for _, feature := range features {
		if !containsInt(featuresIDs, feature.Id) {
			return nil, carError.HandleError(err, "Feature with ID "+strconv.Itoa(feature.Id)+" not found. "+err.Error(), http.StatusNotFound)
		}

	}

	car.Features = features

	var make *mm.Make

	make, err = mr.MakeRepository.FindOne(c.MakeRepository, car.MakeID)

	if err != nil || make == nil {
		return nil, carError.HandleError(err, "Make with ID "+strconv.Itoa(car.MakeID)+" not found. "+err.Error(), http.StatusNotFound)
	}

	var carModel *cmm.CarModel

	carModel, err = cmr.CarModelRepository.FindOne(c.CarModelRepository, car.CarModelID)

	if err != nil || carModel == nil {
		return nil, carError.HandleError(err, "Car Model with ID "+strconv.Itoa(car.CarModel.Id)+" not found. "+err.Error(), http.StatusNotFound)
	}

	var bodyStyle *bsm.BodyStyle

	bodyStyle, err = bsr.BodyStyleRepository.FindOne(c.BodyStyleRepository, car.BodyStyleID)

	if err != nil || bodyStyle == nil {
		return nil, carError.HandleError(err, "Body Style with ID "+strconv.Itoa(car.CarModel.Id)+" not found. "+err.Error(), http.StatusNotFound)
	}

	data, err := c.DBRepository.Transaction(func(i interface{}) (interface{}, error) {
		car, err := c.CarRepository.Create(car)

		return car, err
	})

	car, ok := data.(*model.Car)

	if !ok || err != nil {
		return nil, carError.HandleError(err, "Failed to create Car. "+err.Error(), http.StatusNotModified)
	}

	return car, nil
}

func (ci *carService) Delete(id string) error {
	idn, errValid := strconv.Atoi(id)

	if errValid != nil {
		return carError.HandleError(errValid, "Bad Request. "+errValid.Error(), http.StatusBadRequest)
	}

	errExists := ci.ValidateRecordExists(id)

	if errExists != nil {
		return carError.HandleError(errExists, "Car with ID "+id+" not found. "+errExists.Error(), http.StatusNotFound)
	}

	return ci.CarRepository.Delete(idn)
}

func (ci *carService) Update(car *model.Car) (*model.Car, error) {
	var err error
	var features []*fm.Feature
	var featuresIDs []int

	for _, f := range car.Features {
		featuresIDs = append(featuresIDs, f.Id)
	}

	features, err = fr.FeatureRepository.FindByIDs(ci.FeatureRepository, featuresIDs)

	var idsFound []int

	for _, f := range features {
		idsFound = append(idsFound, f.Id)
	}

	idNotFund, exists := idExists(idsFound, featuresIDs)

	if !exists {
		return nil, carError.HandleError(err, "Feature with ID "+strconv.Itoa(idNotFund)+" not found. ", http.StatusNotFound)
	}

	errExists := ci.ValidateRecordExists(strconv.Itoa(car.Id))

	if errExists != nil {
		return nil, carError.HandleError(errExists, "Car with ID "+strconv.Itoa(car.Id)+" not found. "+errExists.Error(), http.StatusNotFound)
	}

	var make *mm.Make

	make, err = mr.MakeRepository.FindOne(ci.MakeRepository, car.MakeID)

	if err != nil || make == nil {
		return nil, carError.HandleError(err, "Make with ID "+strconv.Itoa(car.MakeID)+" not found. "+err.Error(), http.StatusNotFound)
	}

	car.Make = make

	var carModel *cmm.CarModel

	carModel, err = cmr.CarModelRepository.FindOne(ci.CarModelRepository, car.CarModel.Id)

	if err != nil {
		return nil, carError.HandleError(errExists, "Car Model with ID "+strconv.Itoa(car.CarModel.Id)+" not found. "+err.Error(), http.StatusNotFound)
	}

	var bodyStyle *bsm.BodyStyle

	bodyStyle, err = bsr.BodyStyleRepository.FindOne(ci.BodyStyleRepository, car.BodyStyle.Id)

	if err != nil {
		return nil, carError.HandleError(errExists, "Body Style with ID "+strconv.Itoa(car.CarModel.Id)+" not found. "+err.Error(), http.StatusNotFound)
	}

	var newCarmodel = carModel
	car.CarModel = newCarmodel

	car.BodyStyle = bodyStyle

	data, err := ci.DBRepository.Transaction(func(i interface{}) (interface{}, error) {
		car, err := ci.CarRepository.Update(car)

		return car, err
	})
	car, ok := data.(*model.Car)

	if !ok || err != nil {
		return nil, carError.HandleError(errExists, "Failed to update Car. "+err.Error(), http.StatusNotModified)
	}

	return car, nil
}

func (ci *carService) ValidateRecordExists(id string) error {
	car, err := ci.GetOne(id)

	if err != nil && car == nil {
		return err
	}

	return err
}

func idExists(idsExisting []int, idsSent []int) (int, bool) {
	for _, featureId := range idsSent {
		if !containsInt(idsExisting, featureId) {
			return featureId, false
		}
	}

	return 0, true
}

func containsInt(slice []int, num int) bool {
	for _, v := range slice {
		if v == num {
			return true
		}
	}
	return false
}
