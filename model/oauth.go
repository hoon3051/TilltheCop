package model

import (
	"time"

	"gorm.io/gorm"
)

type Oauth struct {
	gorm.Model
	Provider     string
	AccessToken  string
	RefreshToken string
	Expiry       time.Time

	User_id      uint	`gorm:"unique;not null"`
}
