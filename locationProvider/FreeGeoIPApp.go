package locationprovider

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/spf13/viper"
)

// FreeGeoIPApp geoip provider
type FreeGeoIPApp struct{}

var freeGeoIPAppURL string = "https://freegeoip.app/json/"

type freeGeoIPAppRepsonse struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// GetPublicIPDetails should get the public IP from a service
func (f FreeGeoIPApp) GetPublicIPDetails() error {
	var standardResponse freeGeoIPAppRepsonse

	resp, err := http.Get(freeGeoIPAppURL)
	if err != nil {
		return err
	}

	if resp.StatusCode%200 > 99 {
		return errors.New("Non 200 code returned")
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bodyBytes, &standardResponse)
	if err != nil {
		return err
	}

	viper.Set("latitude", standardResponse.Latitude)
	viper.Set("longitude", standardResponse.Longitude)
	return nil
}
