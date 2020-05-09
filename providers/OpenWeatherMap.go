package provider

import (
	"brick/utils"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/viper"
)

// OpenWeatherMap uses the One Call API to fetch
//   - Current weather
//   - Minute forecast for 1 hour
//   - Hourly forecast for 48 hours
//   - Daily forecast for 7 days
//   - Historical weather data for 5 previous days
type OpenWeatherMap struct {
	Current struct {
		Sunrise    int64   `json:"sunrise"`
		Sunset     int64   `json:"sunset"`
		Temp       float32 `json:"temp"`
		FeelsLike  float32 `json:"feels_like"`
		Pressure   int     `json:"pressure"`
		Humidity   int     `json:"humidity"`
		DewPoint   float32 `json:"dew_point"`
		UVI        float32 `json:"uvi"`
		Clouds     int     `json:"clouds"`
		WindSpeed  float32 `json:"wind_speed"`
		WindDegree float32 `json:"wind_deg"`
		Weather    []struct {
			Main        string `json:"main"`
			Description string `json:"description"`
			Icon        string `json:"icon"`
		} `json:"weather"`
	} `json:"current"`
}

// Using nerd fonts for these icons
var iconCodes = map[string]string{
	"01d": "盛",
	"01n": "",
	"02d": "️",
	"02n": "️",
	"03d": "️",
	"03n": "️",
	"04d": "️",
	"04n": "️",
	"09d": "️",
	"09n": "️",
	"10d": "️",
	"10n": "️",
	"11d": "",
	"11n": "",
	"13d": "️",
	"13n": "",
	"50d": "",
	"50n": "",
}

var owmURL string = "https://api.openweathermap.org/data/2.5/onecall?"

func (o *OpenWeatherMap) populateData() error {
	utils.GetPublicIPDetails()

	var constructedURL string = owmURL + "lat=" + viper.GetString("__BRICK_LAT__") + "&lon=" + viper.GetString("__BRICK_LON__") + "&appid=" + viper.GetString("openweathermap.apikey")
	resp, err := http.Get(constructedURL)
	if err != nil {
		return err
	}

	json.NewDecoder(resp.Body).Decode(&o)
	return nil
}

// CurrentWeather returns a description about the current weather
func (o OpenWeatherMap) CurrentWeather() (string, error) {
	if err := o.populateData(); err != nil {
		return "", err
	}

	var currentDescription string

	if viper.GetBool("__BRICK_TEXTUAL__") {
		currentDescription = o.Current.Weather[0].Main
	} else {
		currentDescription = iconCodes[o.Current.Weather[0].Icon]
	}

	var currentWeather string = fmt.Sprintf("%s %0.2f", currentDescription, o.Current.Temp)
	return currentWeather, nil
}