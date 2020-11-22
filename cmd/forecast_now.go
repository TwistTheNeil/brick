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
		detailed, err := cmd.Flags().GetBool("detailed")
		if err != nil {
			fmt.Println(err)
			return
		}

		if imperial, err = cmd.Flags().GetBool("imperial"); err == nil {
			if textual, err = cmd.Flags().GetBool("textual"); err == nil {
				if detailed {
					if currentWeather, err = p.CurrentWeatherDetailed(imperial, textual); err == nil {
						fmt.Println(currentWeather)
						return
					}
				}
				if currentWeather, err = p.CurrentWeatherShort(imperial, textual); err == nil {
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
	nowCmd.PersistentFlags().Bool("detailed", false, "Show detailed information about weather")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// nowCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// nowCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
