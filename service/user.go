package service

import (
	"errors"

	"github.com/hoon3051/TilltheCop/form"
	"github.com/hoon3051/TilltheCop/model"
	"gorm.io/gorm"
)

type UserService struct{}

func (svc UserService) Register(tx *gorm.DB ,userInfo form.OauthUser) (uint, error) {
	var user model.User
	user.Email = userInfo.Email
	user.OauthId = userInfo.ID
	result := tx.Create(&user)
	if result.Error != nil {
		err := errors.New("failed to register user")
		return user.ID, err
	}

	return user.ID, nil

}
