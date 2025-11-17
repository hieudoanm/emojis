/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"status-cli/services/status"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

// checkCmd represents the status command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Step 1: Choose service
		var service string

		services := make(map[string]string)
		services["cloudflare"] = "https://www.cloudflarestatus.com/api/v2/status.json"
		services["flyio"] = "https://status.flyio.net/api/v2/status.json"
		services["github"] = "https://www.githubstatus.com/api/v2/status.json"
		services["hedera"] = "https://status.hedera.com/api/v2/status.json"
		services["netlify"] = "https://www.netlifystatus.com/api/v2/status.json"
		services["npm"] = "https://status.npmjs.org/api/v2/status.json"
		services["polygon"] = "https://status.polygon.technology/api/v2/status.json"
		services["render"] = "https://status.render.com/api/v2/status.json"
		services["solana"] = "https://status.solana.com/api/v2/status.json"
		services["supabase"] = "https://status.supabase.com/api/v2/status.json"
		services["vercel"] = "https://www.vercel-status.com/api/v2/status.json"

		servicePrompt := &survey.Select{
			Message:  "Choose a service:",
			Options:  []string{"cloudflare", "flyio", "github", "hedera", "netlify", "npm", "polygon", "render", "solana", "supabase", "vercel"},
			Default:  "github",
			PageSize: 1,
		}
		if err := survey.AskOne(servicePrompt, &service); err != nil {
			log.Fatalf("Failed to choose service: %v", err)
		}

		var url string = services[service]
		// Step 2
		status.PrintStatus(url)
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// checkCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// checkCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
