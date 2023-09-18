package controller

import (
	context "api-cars/app/app-common/context"

	"api-cars/app/domain/model"
	"net/http"
)

func SendErrorCar(c context.Context, err error) error {
	e, ok := err.(*model.AppError)

	if e.Code != http.StatusInternalServerError && ok {
		return c.JSON(e.Code, map[string]string{
			"message": e.Message,
		})
	}

	return c.JSON(http.StatusInternalServerError, map[string]string{
		"error": "Internal Server Error",
	})
}
