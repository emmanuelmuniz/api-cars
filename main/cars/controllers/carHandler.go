package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/emmanuelmuniz/api-cars/db"
	"github.com/emmanuelmuniz/api-cars/main/cars/carValidator"
	"github.com/emmanuelmuniz/api-cars/main/cars/models"
	"github.com/emmanuelmuniz/api-cars/main/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

var validate *validator.Validate

func GetAllCars(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var quests []models.Car
	db.DB.Find(&quests)

	json.NewEncoder(w).Encode(quests)
}

func GetCar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]
	var quest models.Car

	if err := db.DB.Where("id = ?", id).First(&quest).Error; err != nil {
		utils.RespondWithError(w, http.StatusNotFound, "Car not found")
		return
	}

	json.NewEncoder(w).Encode(quest)
}

func CreateCar(w http.ResponseWriter, r *http.Request) {
	var car models.Car

	json.NewDecoder(r.Body).Decode(&car)

	createdCar := db.DB.Create(&car)

	err := createdCar.Error

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(car)

}

func UpdateCar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]
	var car models.Car

	if err := db.DB.Where("id = ?", id).First(&car).Error; err != nil {
		utils.RespondWithError(w, http.StatusNotFound, "Car not found")
		return
	}

	var input carValidator.CarInput

	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &input)

	validate = validator.New()
	err := validate.Struct(input)

	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Validation Error")
		return
	}

	car.Make = input.Make
	car.Description = input.Description
	car.Year = input.Year

	db.DB.Save(&car)

	json.NewEncoder(w).Encode(car)
}

func DeleteCar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]
	var car models.Car

	if err := db.DB.Where("id = ?", id).First(&car).Error; err != nil {
		utils.RespondWithError(w, http.StatusNotFound, "Car not found")
		return
	}

	db.DB.Delete(&car)

	w.WriteHeader(http.StatusNoContent)
	json.NewEncoder(w).Encode(car)
}
