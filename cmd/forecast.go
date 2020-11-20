package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// forecastCmd represents the forecast command
var forecastCmd = &cobra.Command{
	Use:   "forecast",
	Short: "Show the weather forecast",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		showText, _ := cmd.Flags().GetBool("textual")
		viper.Set("textual", showText)

		imperial, _ := cmd.Flags().GetBool("imperial")
		if imperial {
			viper.Set("units", "imperial")
			viper.Set("unit_notation", "°F")
		} else {
			viper.Set("units", "metric")
			viper.Set("unit_notation", "°C")
		}
	},
}

func init() {
	rootCmd.AddCommand(forecastCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	forecastCmd.PersistentFlags().Bool("textual", false, "Show textual description instead of icons")
	forecastCmd.PersistentFlags().Bool("imperial", false, "Use imperial system instead of metric")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// forecastCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
