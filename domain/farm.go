package domain

import "gorm.io/gorm"

/*
This is a domain for Farm
*/
type Farm struct {
	gorm.Model
	FarmName string `json:"name"`
	Ponds    []Pond `gorm:"foreignKey:PondID, constraint:OnDelete:CASCADE"`
}
