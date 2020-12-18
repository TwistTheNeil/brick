package locationprovider

import (
	"brick/utils"

	"github.com/spf13/viper"
)

// FreeGeoIPApp geoip provider
type FreeGeoIPApp struct{}

type freeGeoIPAppRepsonse struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// GetPublicIPDetails should get the public IP from a service
func (f FreeGeoIPApp) GetPublicIPDetails() (Location, error) {
	var response freeGeoIPAppRepsonse
	var location Location

	if err := utils.HTTPGet(viper.GetString("locationprovider.url"), &response); err != nil {
		return location, err
	}

	location = Location{
		Latitude:  response.Latitude,
		Longitude: response.Longitude,
	}

	return location, nil
}
