package model

import (
	"time"
)

type Car struct {
	Id          int64      `json:"id"`
	Description string     `json:"description"`
	Make        string     `json:"make"`
	Price       float32    `json:"price"`
	Year        int        `json:"year"`
	New         bool       `json:"new"`
	CarModel    string     `json:"car_model"`
	BodyStyle   string     `json:"body_style"`
	Doors       int        `json:"doors"`
	Distance    int        `json:"distance"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

func (Car) TableName() string { return "cars" }
