package model

import "github.com/jinzhu/gorm"

//Jobs is our job model based on the csv file that was given
type Jobs struct {
	gorm.Model
	Title     string  `json:"title"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}
