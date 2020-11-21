package locationprovider

import (
	"errors"

	"github.com/spf13/viper"
)

// LocationProvider is an interface for
// services providing location data
// such as latitude and longitude
type LocationProvider interface {
	GetPublicIPDetails() error
}

// Selection of location provider
func Selection() (LocationProvider, error) {
	var provider LocationProvider

	switch viper.GetString("locationprovider.name") {
	case "freegeoipapp":
		provider = FreeGeoIPApp{}
	case "restgeoip":
		provider = RestGeoIP{}
	default:
		return provider, errors.New("Unkown location provider")
	}

	return provider, nil
}
