package model

import (
	"time"
)

type Make struct {
	Id          int64      `json:"id"`
	Make        string     `json:"make"`
	Description string     `json:"description"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

func (Make) TableName() string { return "makes" }
