package model

import (
	"time"
)

type BodyStyle struct {
	Id          int        `json:"id"`
	BodyStyle   string     `json:"body_style"`
	Description string     `json:"description"`
	Doors       int        `json:"doors"`
	CreatedAt   *time.Time `json:"-"`
	UpdatedAt   *time.Time `json:"-"`
}

func (BodyStyle) TableName() string { return "body_styles" }
