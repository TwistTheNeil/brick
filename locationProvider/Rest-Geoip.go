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
func (r RestGeoIP) GetPublicIPDetails() error {
	var standardResponse restGeoIPRepsonse

	if err := utils.HTTPGet(viper.GetString("locationprovider.url"), &standardResponse); err != nil {
		return err
	}
	viper.Set("latitude", standardResponse.Location.Latitude)
	viper.Set("longitude", standardResponse.Location.Longitude)
	return nil
}
