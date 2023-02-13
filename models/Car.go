package models

import "gorm.io/gorm"

type Car struct {
	gorm.Model

	Id        string `gorm:"not null; unique_index"`
	Make      string
	Price     float32
	Year      int
	New       bool
	CarModel  string
	BodyStyle string
	Doors     int
	Distance  int
}
