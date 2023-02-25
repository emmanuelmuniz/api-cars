package model

import (
	"time"
)

type Car struct {
	Id          int        `json:"id"`
	MakeID      int        `json:"-"`
	Make        Make       `json:"make" validate:"required"`
	CarModelID  int        `json:"-"`
	CarModel    CarModel   `json:"car_model" validate:"required"`
	Description string     `json:"description"`
	Price       float32    `json:"price" validate:"required"`
	Year        int        `json:"year" validate:"required"`
	New         bool       `json:"new" validate:"required"`
	BodyStyleID int        `json:"-"`
	BodyStyle   BodyStyle  `json:"body_style" validate:"required"`
	Features    []*Feature `json:"features" gorm:"many2many:car_features;"`
	Distance    int        `json:"distance"`
	CreatedAt   *time.Time `json:"-"`
	UpdatedAt   *time.Time `json:"-"`
}

func (Car) TableName() string { return "cars" }
