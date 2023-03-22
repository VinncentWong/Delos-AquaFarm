package domain

import "gorm.io/gorm"

/*
This is a domain for Pond
*/
type Pond struct {
	gorm.Model
	Name   string `json:"name" gorm:"type:varchar(100);not null" validate:"gte=4"`
	Farm   Farm   `validate:"-"`
	FarmID string
}
