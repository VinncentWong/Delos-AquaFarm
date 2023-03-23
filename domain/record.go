package domain

import "gorm.io/gorm"

/*
This is a domain for RecordApi domain
*/
type RecordApi struct {
	gorm.Model
	Endpoint   string `json:"endpoint" gorm:"type:varchar(100)"`
	IpAddress  string `json:"ip_address" gorm:"type:varchar(255)"`
	Count      uint   `json:"count"`
	MethodName string `json:"method_name"`
}
