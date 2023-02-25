package model

import (
	"time"
)

type Feature struct {
	Id          int        `json:"id"`
	Feature     string     `json:"feature"`
	Description string     `json:"description"`
	CreatedAt   *time.Time `json:"-"`
	UpdatedAt   *time.Time `json:"-"`
}

func (Feature) TableName() string { return "features" }
