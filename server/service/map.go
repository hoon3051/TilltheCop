package service

import (
	"fmt"
	"os"

	"github.com/hoon3051/TilltheCop/server/form"
)

type MapService struct{}

func (svc MapService) GetMap(locationForm form.LocationForm) (string, error) {
	// Construct Google Map URL based on latitude and longitude
	googleMapURL := fmt.Sprintf("https://www.google.com/maps/embed/v1/place?key=%s&q=%f,%f&zoom=15", os.Getenv("GOOGLE_MAPS_API_KEY"), locationForm.Location_latitude, locationForm.Location_longitude)
	return googleMapURL, nil
}


