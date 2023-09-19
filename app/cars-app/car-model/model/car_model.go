package model

import (
	m "api-cars/app/cars-app/make/model"

	"time"
)

type CarModel struct {
	Id        int        `json:"id"`
	CarModel  string     `json:"car_model" gorm:"not null" validate:"required"`
	MakeID    int        `json:"make_id" validate:"required"`
	Make      *m.Make    `json:"make" gorm:"foreignKey:MakeID"`
	CreatedAt *time.Time `json:"-"`
	UpdatedAt *time.Time `json:"-"`
}

func (CarModel) TableName() string { return "car_models" }
