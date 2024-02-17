package model

import (
	"gorm.io/gorm"
)

type Record struct {
	gorm.Model
	User_id   uint `gorm:"unique;not null"`
	Report_id uint `gorm:"unique;not null"`
}
