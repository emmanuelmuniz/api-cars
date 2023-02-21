package model

import (
	"time"
)

type Make struct {
	Id          int        `json:"id"`
	Make        string     `json:"make"`
	Description string     `json:"description"`
	CreatedAt   *time.Time `json:"-"`
	UpdatedAt   *time.Time `json:"-"`
}

func (Make) TableName() string { return "makes" }
