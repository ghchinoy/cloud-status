package cmd

import (
	"log"
	"os"

	"github.com/ghchinoy/cloud-status/mcp"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "cloud-status",
	Short: "A CLI to check the status of Google Cloud services.",
	Long:  `A command-line tool to query the Google Cloud status page for current and historical incident data.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// If --mcp or --mcp-http is used, run the MCP server, otherwise show help.
		mcpFlag, _ := cmd.Flags().GetBool("mcp")
		mcpHttpAddr, _ := cmd.Flags().GetString("mcp-http")

		if mcpFlag {
			log.Println("Starting MCP server on stdio...")
			return mcp.Start("")
		}

		if mcpHttpAddr != "" {
			log.Printf("Starting MCP server on HTTP at %s...\n", mcpHttpAddr)
			return mcp.Start(mcpHttpAddr)
		}

		return cmd.Help()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	cobra.OnInitialize(initConfig)

	rootCmd.Flags().Bool("mcp", false, "Start the application in MCP server mode over stdio.")
	rootCmd.Flags().String("mcp-http", "", "Start the application in MCP server mode over HTTP, listening on the given address.")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	viper.AutomaticEnv() // read in environment variables that match
}
