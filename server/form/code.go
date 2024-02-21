package form

import (
	"github.com/go-playground/validator/v10"
)

type CodeData struct {
	ReportID string `json:"reportID" binding:"required"`
}

func (f CodeData) Validate() string {
	validate := validator.New()
	err := validate.Struct(f)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return err.Field() + " is required"
		}
	}
	return ""
}