package model

import (
	"time"
)

type CarImage struct {
	Id          int        `json:"id"`
	CarID       int        `json:"car_id"`
	ImageURL    string     `json:"image_url" gorm:"not null" validate:"required"`
	Description string     `json:"description" gorm:"not null" validate:"required"`
	CreatedAt   *time.Time `json:"-"`
	UpdatedAt   *time.Time `json:"-"`
}

func (CarImage) TableName() string { return "car_images" }
