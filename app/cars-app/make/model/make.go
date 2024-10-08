package model

import (
	"time"
)

type Make struct {
	Id          int        `json:"id"`
	Make        string     `json:"make" gorm:"not null" validate:"required"`
	Description string     `json:"description" gorm:"not null" validate:"required"`
	CreatedAt   *time.Time `json:"-"`
	UpdatedAt   *time.Time `json:"-"`
}

func (Make) TableName() string { return "makes" }
