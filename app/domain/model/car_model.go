package model

import (
	"time"
)

type CarModel struct {
	Id        int64      `json:"id"`
	MakeID    int64      `json:"-"`
	CarModel  string     `json:"car_model"`
	Make      Make       `json:"make" gorm:"foreignKey:MakeID;references:id"`
	ModelMake string     `json:"model_make"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

func (CarModel) TableName() string { return "car_models" }
