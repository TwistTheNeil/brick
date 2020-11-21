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
func (f FreeGeoIPApp) GetPublicIPDetails() error {
	var standardResponse freeGeoIPAppRepsonse

	if err := utils.HTTPGet(viper.GetString("locationprovider.url"), &standardResponse); err != nil {
		return err
	}
	viper.Set("latitude", standardResponse.Latitude)
	viper.Set("longitude", standardResponse.Longitude)
	return nil
}
