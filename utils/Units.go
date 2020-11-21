package utils

// UnitNotation displays the notation for temperature units
func UnitNotation(units string) string {
	notation := ""

	switch units {
	case "metric":
		notation = "°C"
	case "imperial":
		notation = "°F"
	default:
		notation = units
	}

	return notation
}
