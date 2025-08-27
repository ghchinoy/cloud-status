package cmd

import (
	"fmt"
	"github.com/ghchinoy/cloud-status/internal/fetcher"
	"github.com/ghchinoy/cloud-status/internal/parser"
	"github.com/spf13/cobra"
)

// currentCmd represents the current command
var currentCmd = &cobra.Command{
	Use:   "current",
	Short: "Get the current status from the Google Cloud Atom feed.",
	Long:  `Fetches and displays the most recent incident updates from the Google Cloud status Atom feed.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Fetching current status...")

		atomData, err := fetcher.FetchFeedAtom()
		if err != nil {
			return fmt.Errorf("could not fetch atom feed: %w", err)
		}

		feed, err := parser.ParseAtomFeed(atomData)
		if err != nil {
			return fmt.Errorf("could not parse atom feed: %w", err)
		}

		if len(feed.Entries) == 0 {
			fmt.Println("No recent incidents found in the feed.")
			return nil
		}

		fmt.Println("\n--- Most Recent Google Cloud Status Updates ---")
		for _, entry := range feed.Entries {
			fmt.Printf("\nTitle: %s\n", entry.Title)
			fmt.Printf("Updated: %s\n", entry.Updated)
			fmt.Printf("ID: %s\n", entry.ID)
			// The content is HTML, so we are not printing it directly for now.
			// In the future, we could strip the HTML tags.
		}
		fmt.Println("\n---------------------------------------------")

		return nil
	},
}

func init() {
	rootCmd.AddCommand(currentCmd)
}
