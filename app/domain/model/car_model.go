package model

import (
	"time"
)

type CarModel struct {
	Id        int        `json:"id"`
	CarModel  string     `json:"car_model" gorm:"not null"`
	MakeID    int        `json:"-"`
	Make      Make       `json:"make" gorm:"not null"`
	CreatedAt *time.Time `json:"-"`
	UpdatedAt *time.Time `json:"-"`
}

func (CarModel) TableName() string { return "car_models" }
