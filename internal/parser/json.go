package parser

import (
	"encoding/json"
	"github.com/ghchinoy/cloud-status/internal/types"
)

// ParseIncidents parses the JSON data into a slice of Incident structs.
func ParseIncidents(data []byte) ([]types.Incident, error) {
	var incidents []types.Incident
	err := json.Unmarshal(data, &incidents)
	if err != nil {
		return nil, err
	}
	return incidents, nil
}
