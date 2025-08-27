package types

import "time"

// Incident represents a single incident from the incidents.json feed.
type Incident struct {
	ID                string    `json:"id"`
	Begin             time.Time `json:"begin"`
	End               time.Time `json:"end"`
	IncidentID        string    `json:"incident_id"`
	ServiceKey        string    `json:"service_key"`
	ServiceName       string    `json:"service_name"`
	Severity          string    `json:"severity"`
	URI               string    `json:"uri"`
	Created           time.Time `json:"created"`
	Modified          time.Time `json:"modified"`
	MostRecentUpdate  Update    `json:"most_recent_update"`
	Updates           []Update  `json:"updates"`
}

// Update represents a single update within an incident.
type Update struct {
	Created  time.Time `json:"created"`
	Modified time.Time `json:"modified"`
	When     time.Time `json:"when"`
	Text     string    `json:"text"`
	Status   string    `json:"status"`
}

// AtomFeed represents the structure of the feed.atom file.
type AtomFeed struct {
	Entries []AtomEntry `xml:"entry"`
}

// AtomEntry represents a single entry in the Atom feed.
type AtomEntry struct {
	ID      string `xml:"id"`
	Title   string `xml:"title"`
	Updated string `xml:"updated"`
	Content string `xml:"content"`
}
