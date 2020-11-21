package cmd

import (
	"github.com/spf13/cobra"
)

// forecastCmd represents the forecast command
var forecastCmd = &cobra.Command{
	Use:   "forecast",
	Short: "Show the weather forecast",
}

func init() {
	rootCmd.AddCommand(forecastCmd)
}
