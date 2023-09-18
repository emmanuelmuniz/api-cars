package model

import (
	bs "api-cars/app/cars-app/body-style/model"
	cm "api-cars/app/cars-app/car-model/model"
	f "api-cars/app/cars-app/feature/model"
	m "api-cars/app/cars-app/make/model"

	"time"
)

type Car struct {
	Id          int           `json:"id"`
	MakeID      int           `json:"-"`
	Make        *m.Make       `json:"make" validate:"required"`
	CarModelID  int           `json:"-"`
	CarModel    *cm.CarModel  `json:"car_model" validate:"required"`
	Description string        `json:"description"`
	Price       float32       `json:"price" validate:"required"`
	Year        int           `json:"year" validate:"required"`
	New         bool          `json:"new" validate:"required"`
	BodyStyleID int           `json:"-"`
	BodyStyle   *bs.BodyStyle `json:"body_style" validate:"required"`
	Features    []*f.Feature  `json:"features" gorm:"many2many:car_features;" validate:"required"`
	Distance    int           `json:"distance"`
	CreatedAt   *time.Time    `json:"-"`
	UpdatedAt   *time.Time    `json:"-"`
}

func (Car) TableName() string { return "cars" }
