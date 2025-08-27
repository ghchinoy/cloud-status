package cmd

import (
	"fmt"
	"strings"

	"github.com/ghchinoy/cloud-status/internal/fetcher"
	"github.com/ghchinoy/cloud-status/internal/parser"
	"github.com/ghchinoy/cloud-status/internal/types"
	"github.com/spf13/cobra"
)

// historyCmd represents the history command
var historyCmd = &cobra.Command{
	Use:   "history",
	Short: "Get historical incident data from the Google Cloud JSON feed.",
	Long:  `Fetches and displays historical incident data from the Google Cloud status JSON feed, with options to filter.`, 
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Fetching historical incidents...")

		jsonData, err := fetcher.FetchIncidentsJSON()
		if err != nil {
			return fmt.Errorf("could not fetch incidents json: %w", err)
		}

		incidents, err := parser.ParseIncidents(jsonData)
		if err != nil {
			return fmt.Errorf("could not parse incidents json: %w", err)
		}

		serviceFilter, _ := cmd.Flags().GetString("service")
		severityFilter, _ := cmd.Flags().GetString("severity")
		limit, _ := cmd.Flags().GetInt("limit")

		var filteredIncidents []types.Incident

		for _, incident := range incidents {
			if serviceFilter != "" && !strings.EqualFold(incident.ServiceName, serviceFilter) {
				continue
			}
			if severityFilter != "" && !strings.EqualFold(incident.Severity, severityFilter) {
				continue
			}
			filteredIncidents = append(filteredIncidents, incident)
		}

		if len(filteredIncidents) == 0 {
			fmt.Println("No incidents found matching the criteria.")
			return nil
		}

		// Apply limit
		if limit > 0 && len(filteredIncidents) > limit {
			filteredIncidents = filteredIncidents[:limit]
		}

		fmt.Printf("\n--- Found %d Incidents ---\n", len(filteredIncidents))
		for _, incident := range filteredIncidents {
			fmt.Printf("\nService: %s\n", incident.ServiceName)
			fmt.Printf("Severity: %s\n", incident.Severity)
			fmt.Printf("Began: %s\n", incident.Begin.Format("2006-01-02 15:04 MST"))
			fmt.Printf("Ended: %s\n", incident.End.Format("2006-01-02 15:04 MST"))
			fmt.Printf("URL: https://status.cloud.google.com%s\n", incident.URI)
			fmt.Printf("Most Recent Update: %s\n", incident.MostRecentUpdate.Text)
		}
		fmt.Println("\n-------------------------")

		return nil
	},
}

func init() {
	rootCmd.AddCommand(historyCmd)

	historyCmd.Flags().String("service", "", "Filter incidents by service name (e.g., 'Google Compute Engine')")
	historyCmd.Flags().String("severity", "", "Filter incidents by severity (low, medium, high)")
	historyCmd.Flags().Int("limit", 10, "Limit the number of incidents returned")
}
