package form

import (
	"github.com/go-playground/validator/v10"
)

type LocationForm struct {
	Location_latitude float64 `json:"location_latitude" binding:"required"`
	Location_longitude float64 `json:"location_longitude" binding:"required"`
}

func (f LocationForm) Validate() string{
	validate := validator.New()
	err := validate.Struct(f)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return err.Field() + " is required"
		}
	}
	return ""
}