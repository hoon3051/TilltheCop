package model

import (
	"gorm.io/gorm"
)

type Profile struct {
	gorm.Model
	Name 	string
	Age 	int
	Gender 	string

	User_id uint	`gorm:"unique;not null"`
}