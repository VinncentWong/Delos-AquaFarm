package domain

import "gorm.io/gorm"

/*
This is a domain for Farm
*/
type Farm struct {
	gorm.Model
	FarmName string `json:"name" binding:"required" gorm:"type:varchar(50)"`
	Ponds    []Pond `gorm:"foreignKey:FarmID;constraint:OnDelete:CASCADE" json:"-"`
}
