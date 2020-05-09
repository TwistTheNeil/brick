package cmd

import (
	provider "brick/providers"
	"fmt"

	"github.com/spf13/cobra"
)

// nowCmd represents the now command
var nowCmd = &cobra.Command{
	Use:   "now",
	Short: "Show the current weather forecast",
	Run: func(cmd *cobra.Command, args []string) {
		var p provider.OpenWeatherMap
		currentWeather, _ := p.CurrentWeather()
		fmt.Println(currentWeather)
	},
}

func init() {
	forecastCmd.AddCommand(nowCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// nowCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// nowCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
