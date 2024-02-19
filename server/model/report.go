package model

import (
	"gorm.io/gorm"
)

type Report struct {
	gorm.Model
	Crime 	string
	Location_latitude 	string
	Location_longitude 	string

	User_id uint	`gorm:"unique;not null"`
}