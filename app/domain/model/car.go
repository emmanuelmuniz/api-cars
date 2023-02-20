package model

import (
	"time"
)

type Car struct {
	Id          int64      `json:"id"`
	MakeID      int64      `json:"-"`
	Make        Make       `json:"make" gorm:"foreignKey:MakeID;references:id"`
	CarModelID  int64      `json:"-"`
	CarModel    CarModel   `json:"car_models" gorm:"foreignKey:CarModelID;references:id"`
	Description string     `json:"description"`
	Price       float32    `json:"price"`
	Year        int        `json:"year"`
	New         bool       `json:"new"`
	BodyStyle   string     `json:"body_style"`
	Doors       int        `json:"doors"`
	Distance    int        `json:"distance"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

func (Car) TableName() string { return "cars" }
