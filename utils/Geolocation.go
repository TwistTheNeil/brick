package utils

import (
	"encoding/json"
	"net/http"

	"github.com/spf13/viper"
)

// FreeGeoIPAppResponse defines the GET response from freegeoip.app
type FreeGeoIPAppResponse struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

var freeGeoIPAppURL string = "https://freegeoip.app/json/"

// GetPublicIPDetails should get the public IP from a service
func GetPublicIPDetails() error {
	var g FreeGeoIPAppResponse

	resp, err := http.Get(freeGeoIPAppURL)
	if err != nil {
		return err
	}

	json.NewDecoder(resp.Body).Decode(&g)

	viper.Set("latitude", g.Latitude)
	viper.Set("longitude", g.Longitude)
	return nil
}
