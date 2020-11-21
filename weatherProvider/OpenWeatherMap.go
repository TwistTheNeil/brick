package weatherprovider

import (
	locationprovider "brick/locationProvider"
	"brick/utils"
	"fmt"

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

func (o *OpenWeatherMap) populateData(units string) error {
	locationProvider, err := locationprovider.Selection()
	if err != nil {
		return err
	}

	location, err := locationProvider.GetPublicIPDetails()
	if err != nil {
		return err
	}

	var constructedURL string = owmURL + "lat=" + fmt.Sprintf("%.4f", location.Latitude) + "&lon=" + fmt.Sprintf("%.4f", location.Longitude) + "&units=" + units + "&appid=" + viper.GetString("openweathermap.apikey")

	err = utils.HTTPGet(constructedURL, &o)
	if err != nil {
		return err
	}

	return nil
}

// CurrentWeather returns a description about the current weather
func (o OpenWeatherMap) CurrentWeather(imperialSystem, textualOutput bool) (string, error) {
	units := "metric"
	if imperialSystem {
		units = "imperial"
	}

	if err := o.populateData(units); err != nil {
		return "", err
	}

	var currentDescription string

	if textualOutput {
		currentDescription = o.Current.Weather[0].Main
	} else {
		currentDescription = iconCodes[o.Current.Weather[0].Icon]
	}

	var currentWeather string = fmt.Sprintf("%s %0.2f%s", currentDescription, o.Current.Temp, utils.UnitNotation(units))
	return currentWeather, nil
}
