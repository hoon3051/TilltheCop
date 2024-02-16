package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email	string `gorm:"unique"`
	OauthId	string `gorm:"unique"`

	Oauth Oauth `gorm:"foreignKey:User_id"`
}
