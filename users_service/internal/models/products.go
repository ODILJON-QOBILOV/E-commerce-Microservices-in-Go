package models

import "gorm.io/gorm"

type Products struct {
	gorm.Model
	Title       string
	Description string
	Price       int
	Amount      int
	Image       string
	
	CategoryId  int
	Category    Category
}
