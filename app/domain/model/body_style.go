package model

import (
	"time"
)

type BodyStyle struct {
	Id          int        `json:"id"`
	BodyStyle   string     `json:"body_style" gorm:"not null"`
	Description string     `json:"description"`
	Doors       int        `json:"doors" gorm:"not null"`
	CreatedAt   *time.Time `json:"-"`
	UpdatedAt   *time.Time `json:"-"`
}

func (BodyStyle) TableName() string { return "body_styles" }
