package mcp

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/ghchinoy/cloud-status/internal/fetcher"
	"github.com/ghchinoy/cloud-status/internal/parser"
	"github.com/ghchinoy/cloud-status/internal/types"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// GetCloudStatusArgs defines the arguments for the get_cloud_status tool.
type GetCloudStatusArgs struct {
	Source        string `json:"source"` // "current" or "history"
	ServiceFilter string `json:"service_filter,omitempty"`
	SeverityFilter string `json:"severity_filter,omitempty"`
	Limit         int    `json:"limit,omitempty"`
}

// Start starts the MCP server.
func Start(httpAddr string) error {
	server := mcp.NewServer(&mcp.Implementation{Name: "cloud-status"}, nil)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_cloud_status",
		Description: "Retrieves status information for Google Cloud services. It can get the current status from the Atom feed or historical incidents from the JSON feed. Parameters: 'source' (required: 'current' or 'history'), 'service_filter' (optional, for history), 'severity_filter' (optional, for history), 'limit' (optional, for history).",
	}, getCloudStatusHandler)

	if httpAddr != "" {
		handler := mcp.NewStreamableHTTPHandler(func(*http.Request) *mcp.Server {
			return server
		}, nil)
		log.Printf("MCP handler listening at %s", httpAddr)
		return http.ListenAndServe(httpAddr, handler)
	}

	// MCP over stdio
	t := mcp.NewStdioTransport()
	if err := server.Run(context.Background(), t); err != nil {
		log.Printf("Server failed: %v", err)
		return err
	}

	return nil
}

// getCloudStatusHandler is the handler for the get_cloud_status tool.
func getCloudStatusHandler(ctx context.Context, req *mcp.CallToolRequest, args GetCloudStatusArgs) (*mcp.CallToolResult, any, error) {
	source := strings.ToLower(args.Source)
	if source == "" {
		return nil, nil, fmt.Errorf("source argument is required (either 'current' or 'history')")
	}

	var output string

	if source == "current" {
		atomData, err := fetcher.FetchFeedAtom()
		if err != nil {
			return nil, nil, fmt.Errorf("could not fetch atom feed: %w", err)
		}
		feed, err := parser.ParseAtomFeed(atomData)
		if err != nil {
			return nil, nil, fmt.Errorf("could not parse atom feed: %w", err)
		}
		if len(feed.Entries) == 0 {
			output = "No recent incidents found in the feed."
		} else {
			var builder strings.Builder
			builder.WriteString("--- Most Recent Google Cloud Status Updates ---\n")
			for _, entry := range feed.Entries {
				builder.WriteString(fmt.Sprintf("\nTitle: %s\n", entry.Title))
				builder.WriteString(fmt.Sprintf("Updated: %s\n", entry.Updated))
			}
			output = builder.String()
		}
	} else if source == "history" {
		jsonData, err := fetcher.FetchIncidentsJSON()
		if err != nil {
			return nil, nil, fmt.Errorf("could not fetch incidents json: %w", err)
		}
		incidents, err := parser.ParseIncidents(jsonData)
		if err != nil {
			return nil, nil, fmt.Errorf("could not parse incidents json: %w", err)
		}

		var filteredIncidents []types.Incident
		for _, incident := range incidents {
			if args.ServiceFilter != "" && !strings.EqualFold(incident.ServiceName, args.ServiceFilter) {
				continue
			}
			if args.SeverityFilter != "" && !strings.EqualFold(incident.Severity, args.SeverityFilter) {
				continue
			}
			filteredIncidents = append(filteredIncidents, incident)
		}

		if len(filteredIncidents) == 0 {
			output = "No incidents found matching the criteria."
		} else {
			limit := args.Limit
			if limit > 0 && len(filteredIncidents) > limit {
				filteredIncidents = filteredIncidents[:limit]
			}
			var builder strings.Builder
			builder.WriteString(fmt.Sprintf("--- Found %d Incidents ---\n", len(filteredIncidents)))
			for _, incident := range filteredIncidents {
				builder.WriteString(fmt.Sprintf("\nService: %s\n", incident.ServiceName))
				builder.WriteString(fmt.Sprintf("Severity: %s\n", incident.Severity))
				builder.WriteString(fmt.Sprintf("Began: %s\n", incident.Begin.Format("2006-01-02 15:04 MST")))
				builder.WriteString(fmt.Sprintf("URL: https://status.cloud.google.com%s\n", incident.URI))
			}
			output = builder.String()
		}
	} else {
		return nil, nil, fmt.Errorf("invalid source: %s. Must be 'current' or 'history'", source)
	}

	result := &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: output},
		},
	}
	return result, nil, nil
}
