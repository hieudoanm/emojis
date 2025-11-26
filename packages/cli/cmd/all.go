/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"status-cli/configs"
	"status-cli/services/status"

	"github.com/spf13/cobra"
)

var debug bool // debug flag

// allCmd represents the "all" status command
var allCmd = &cobra.Command{
	Use:   "all",
	Short: "Show status of all services",
	Long:  `Show the current status for all configured services, optionally with debug logging.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Iterate over all services and show their status
		for name, url := range configs.Services {
			// Pass the debug flag into the requests.Options
			status.PrintDescriptiveStatus(name, url, debug)
		}
	},
}

func init() {
	rootCmd.AddCommand(allCmd)

	// Add a local flag --debug to enable verbose debug logging
	allCmd.Flags().BoolVarP(&debug, "debug", "d", false, "Enable debug logging for HTTP requests")
}
