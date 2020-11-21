package weatherprovider

// Provider is an interface for any weather provider
type Provider interface {
	CurrentWeather() string
}
