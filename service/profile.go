package service

import (
	"errors"

	"github.com/hoon3051/TilltheCop/form"
	"github.com/hoon3051/TilltheCop/initializer"
	"github.com/hoon3051/TilltheCop/model"
	"gorm.io/gorm"
)

type ProfileService struct{}

func (svc ProfileService) CreateProfile(tx *gorm.DB , profileform form.ProfileForm, userID uint) (error) {
	var profile model.Profile
	profile.Name = profileform.Name
	profile.Age = profileform.Age
	profile.Gender = profileform.Gender
	profile.User_id = userID
	result := tx.Create(&profile)
	if result.Error != nil {
		err := errors.New("failed to create profile")
		return err
	}

	return nil

}

func (svc ProfileService) GetProfile(userID uint) (model.Profile, error) {
	var profile model.Profile
	result := initializer.DB.First(&profile, "user_id = ?", userID)
	if result.Error != nil {
		err := errors.New("failed to get profile")
		return profile, err
	}

	return profile, nil
}

func (svc ProfileService) UpdateProfile(userID uint, profileform form.ProfileForm) (model.Profile, error) {
	var profile model.Profile
	result := initializer.DB.First(&profile, "user_id = ?", userID)
	if result.Error != nil {
		err := errors.New("failed to get profile")
		return profile, err
	}
	profile.Name = profileform.Name
	profile.Age = profileform.Age
	profile.Gender = profileform.Gender
	result = initializer.DB.Save(&profile)
	if result.Error != nil {
		err := errors.New("failed to update profile")
		return profile, err
	}

	return profile, nil

}

