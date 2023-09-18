package validator

import (
	"api-cars/app/domain/model"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func ValidateStruct(s interface{}) error {
	validate := validator.New()
	err := validate.Struct(s)

	if err != nil {
		return model.HandleError(err, err.Error(), http.StatusBadRequest)
	}

	return err
}
