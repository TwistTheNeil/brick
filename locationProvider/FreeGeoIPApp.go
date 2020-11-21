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
func (f FreeGeoIPApp) GetPublicIPDetails() (FlattenedProviderResponse, error) {
	var standardResponse freeGeoIPAppRepsonse
	var returnResponse FlattenedProviderResponse

	if err := utils.HTTPGet(viper.GetString("locationprovider.url"), &standardResponse); err != nil {
		return returnResponse, err
	}

	returnResponse = FlattenedProviderResponse{
		Latitude:  standardResponse.Latitude,
		Longitude: standardResponse.Longitude,
	}

	return returnResponse, nil
}
