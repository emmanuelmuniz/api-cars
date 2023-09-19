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
	Name        string        `json:"name" validate:"required"`
	MakeID      int           `json:"make_id" validate:"required"`
	Make        *m.Make       `json:"make" gorm:"foreignKey:MakeID"`
	CarModelID  int           `json:"car_model_id" validate:"required"`
	CarModel    *cm.CarModel  `json:"car_model" gorm:"foreignKey:CarModelID"`
	Description string        `json:"description"`
	Price       float32       `json:"price" validate:"required"`
	Year        int           `json:"year" validate:"required"`
	New         bool          `json:"new" validate:"required"`
	BodyStyleID int           `json:"body_style_id" validate:"required"`
	BodyStyle   *bs.BodyStyle `json:"body_style" gorm:"foreignKey:BodyStyleID"`
	Features    []*f.Feature  `json:"features" gorm:"many2many:car_features;" validate:"required"`
	Images      []string      `json:"images" gorm:"type:varchar[]"`
	Distance    int           `json:"distance"`
	CreatedAt   *time.Time    `json:"-"`
	UpdatedAt   *time.Time    `json:"-"`
}

func (Car) TableName() string { return "cars" }
