package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email	string `gorm:"unique"`
	OauthId	string `gorm:"unique"`

	Oauth Oauth `gorm:"foreignKey:User_id"`
	Profile Profile `gorm:"foreignKey:User_id"`
	Reports []Report `gorm:"foreignKey:User_id"`
	Records []Record `gorm:"foreignKey:User_id"`
}
