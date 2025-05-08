package models

import "gorm.io/gorm"

type Tast struct {
	gorm.Model
	Title string `json:"title"`
	Description string `json:"description"`
	Completed bool `json:"completed"`
}