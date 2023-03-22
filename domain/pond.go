package domain

import "gorm.io/gorm"

/*
This is a domain for Pond
*/
type Pond struct {
	gorm.Model
	Name   string `json:"name" gorm:"type:varchar(100)" binding:"required"`
	Farm   Farm
	FarmID uint
}
