package model

import (
	"time"
)

type CarModel struct {
	Id        int        `json:"id"`
	MakeID    int        `json:"-"`
	Make      Make       `json:"make"`
	CarModel  string     `json:"car_model"`
	CreatedAt *time.Time `json:"-"`
	UpdatedAt *time.Time `json:"-"`
}

func (CarModel) TableName() string { return "car_models" }
