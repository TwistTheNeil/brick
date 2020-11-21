package weatherprovider

// Provider is an interface for any weather provider
type Provider interface {
	// Arguments:
	//   - imperialSystem
	//   - textual
	CurrentWeather(bool, bool) string
}
