package model

import (
	"time"
)

type CarModel struct {
	Id        int        `json:"id"`
	CarModel  string     `json:"car_model"`
	MakeID    int        `json:"-"`
	Make      Make       `json:"make"`
	CreatedAt *time.Time `json:"-"`
	UpdatedAt *time.Time `json:"-"`
}

func (CarModel) TableName() string { return "car_models" }
