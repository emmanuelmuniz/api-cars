package model

import (
	"time"
)

type Car struct {
	Id          int        `json:"id"`
	MakeID      int        `json:"-"`
	Make        Make       `json:"make"`
	CarModelID  int        `json:"-"`
	CarModel    CarModel   `json:"car_model"`
	Description string     `json:"description"`
	Price       float32    `json:"price"`
	Year        int        `json:"year"`
	New         bool       `json:"new"`
	BodyStyleID int        `json:"-"`
	BodyStyle   BodyStyle  `json:"body_style"`
	Features    []*Feature `json:"features" gorm:"many2many:car_features;"`
	Doors       int        `json:"doors"`
	Distance    int        `json:"distance"`
	CreatedAt   *time.Time `json:"-"`
	UpdatedAt   *time.Time `json:"-"`
}

func (Car) TableName() string { return "cars" }
