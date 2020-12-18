package locationprovider

import (
	"brick/utils"

	"github.com/spf13/viper"
)

// Maxmind service
type Maxmind struct{}

type maxmindRepsonse struct {
	Location struct {
		Latitude  float64 `json:"Latitude"`
		Longitude float64 `json:"Longitude"`
	} `json:"Location"`
}

// GetPublicIPDetails should get the public IP from a service
func (m Maxmind) GetPublicIPDetails() (Location, error) {
	var response maxmindRepsonse
	var location Location

	if err := utils.HTTPGet(viper.GetString("locationprovider.url"), &response); err != nil {
		return location, err
	}

	location = Location{
		Latitude:  response.Location.Latitude,
		Longitude: response.Location.Longitude,
	}

	return location, nil
}
