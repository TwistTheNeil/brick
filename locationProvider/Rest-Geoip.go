package locationprovider

import (
	"brick/utils"

	"github.com/spf13/viper"
)

// RestGeoIP geoip provider
type RestGeoIP struct{}

type restGeoIPRepsonse struct {
	Location struct {
		Latitude  float64 `json:"Latitude"`
		Longitude float64 `json:"Longitude"`
	} `json:"Location"`
}

// GetPublicIPDetails should get the public IP from a service
func (r RestGeoIP) GetPublicIPDetails() (FlattenedProviderResponse, error) {
	var standardResponse restGeoIPRepsonse
	var returnResponse FlattenedProviderResponse

	if err := utils.HTTPGet(viper.GetString("locationprovider.url"), &standardResponse); err != nil {
		return returnResponse, err
	}

	returnResponse = FlattenedProviderResponse{
		Latitude:  standardResponse.Location.Latitude,
		Longitude: standardResponse.Location.Longitude,
	}

	return returnResponse, nil
}
