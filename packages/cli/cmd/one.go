/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"status-cli/configs"
	"status-cli/services/status"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

// oneCmd represents the "one" status command
var oneCmd = &cobra.Command{
	Use:   "one",
	Short: "Show status of a single service",
	Long:  `Prompt the user to select a service and show its full status, optionally with debug logging.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Build the service options
		options := make([]string, 0, len(configs.Services))
		for k := range configs.Services {
			options = append(options, k)
		}

		// Step 1: Choose service
		var service string
		servicePrompt := &survey.Select{
			Message:  "Choose a service:",
			Options:  options,
			Default:  "github",
			PageSize: 5,
		}
		if err := survey.AskOne(servicePrompt, &service); err != nil {
			log.Fatalf("Failed to choose service: %v", err)
		}

		// Step 2: Fetch and display status with debug flag
		url := configs.Services[service]
		status.PrintFullStatus(url, debug)
	},
}

func init() {
	rootCmd.AddCommand(oneCmd)

	// Add a local --debug flag for this command
	oneCmd.Flags().BoolVarP(&debug, "debug", "d", false, "Enable debug logging for HTTP requests")
}
