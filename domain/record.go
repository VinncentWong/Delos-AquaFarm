package domain

import "gorm.io/gorm"

/*
This is a domain for RecordApi domain
*/
type RecordApi struct {
	gorm.Model
	Endpoint  string
	IpAddress string
	Count     uint
}
