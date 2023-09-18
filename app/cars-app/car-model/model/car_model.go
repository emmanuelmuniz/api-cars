package model

import (
	m "api-cars/app/cars-app/make/model"

	"time"
)

type CarModel struct {
	Id        int        `json:"id"`
	CarModel  string     `json:"car_model" gorm:"not null" validate:"required"`
	MakeID    int        `json:"-"`
	Make      *m.Make    `json:"make" gorm:"not null" validate:"required"`
	CreatedAt *time.Time `json:"-"`
	UpdatedAt *time.Time `json:"-"`
}

func (CarModel) TableName() string { return "car_models" }
