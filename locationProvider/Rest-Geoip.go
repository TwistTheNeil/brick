package locationprovider

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/spf13/viper"
)

// RestGeoIP geoip provider
type RestGeoIP struct{}

var url string = "https://ip.neilcastelino.com/api/geoip"

type restGeoIPRepsonse struct {
	Location struct {
		Latitude  float64 `json:"Latitude"`
		Longitude float64 `json:"Longitude"`
	} `json:"Location"`
}

// GetPublicIPDetails should get the public IP from a service
func (r RestGeoIP) GetPublicIPDetails() error {
	var standardResponse restGeoIPRepsonse

	resp, err := http.Get(url)
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

	viper.Set("latitude", standardResponse.Location.Latitude)
	viper.Set("longitude", standardResponse.Location.Longitude)
	return nil
}
