package model

import "github.com/jinzhu/gorm"

type details struct {
	gorm.Model
	paitentName string
	age         int
	mnumber     int
}
