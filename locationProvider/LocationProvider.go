package location

import (
	"errors"

	"github.com/spf13/viper"
)

// Provider is an interface for
// services providing location data
// such as latitude and longitude
type Provider interface {
	GetPublicIPDetails() (Location, error)
}

// Location represents map coordinates
type Location struct {
	Latitude  float64
	Longitude float64
}

// Selection of location provider
func Selection() (Provider, error) {
	var provider Provider

	switch viper.GetString("locationprovider.name") {
	case "freegeoipapp":
		provider = FreeGeoIPApp{}
	case "maxmind":
		provider = Maxmind{}
	default:
		return provider, errors.New("Unkown location provider")
	}

	return provider, nil
}
