package model

import (
	"time"
)

type Feature struct {
	Id          int        `json:"id"`
	Feature     string     `json:"feature" gorm:"not null" validate:"required"`
	Description string     `json:"description" gorm:"not null" validate:"required"`
	CreatedAt   *time.Time `json:"-"`
	UpdatedAt   *time.Time `json:"-"`
}

func (Feature) TableName() string { return "features" }
