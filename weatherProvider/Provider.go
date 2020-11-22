package weatherprovider

// Provider is an interface for any weather provider
type Provider interface {
	// Arguments:
	//   - imperialSystem
	//   - textual
	CurrentWeatherShort(bool, bool) string
	CurrentWeatherDetailed(bool, bool) string
}
