package cmd

import (
	weatherprovider "brick/weatherProvider"
	"fmt"

	"github.com/spf13/cobra"
)

// nowCmd represents the now command
var nowCmd = &cobra.Command{
	Use:   "now",
	Short: "Show the current weather forecast",
	Run: func(cmd *cobra.Command, args []string) {
		var p weatherprovider.OpenWeatherMap
		var err error
		var imperial, textual bool
		var currentWeather string

		if imperial, err = cmd.Flags().GetBool("imperial"); err == nil {
			if textual, err = cmd.Flags().GetBool("textual"); err == nil {
				if currentWeather, err = p.CurrentWeather(imperial, textual); err == nil {
					fmt.Println(currentWeather)
				}
			}
		}

		if err != nil {
			fmt.Println(err)
		}
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
