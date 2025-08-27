package parser

import (
	"encoding/xml"
	"github.com/ghchinoy/cloud-status/internal/types"
)

// ParseAtomFeed parses the Atom feed XML data into an AtomFeed struct.
func ParseAtomFeed(data []byte) (*types.AtomFeed, error) {
	var feed types.AtomFeed
	err := xml.Unmarshal(data, &feed)
	if err != nil {
		return nil, err
	}
	return &feed, nil
}
