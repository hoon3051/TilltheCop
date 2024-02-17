package form

import (
	"github.com/go-playground/validator/v10"
)


type ProfileForm struct {
	Name 	string `json:"name"`
	Age 	int `json:"age"`
	Gender 	string `json:"gender"`
}

func (f ProfileForm) Validate() string{
	validate := validator.New()
	err := validate.Struct(f)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return err.Field() + " is required"
		}
	}
	return ""
}