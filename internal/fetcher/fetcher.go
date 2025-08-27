package fetcher

import (
	"fmt"
	"io"
	"net/http"
)

const (
	jsonURL = "https://status.cloud.google.com/incidents.json"
	atomURL = "https://status.cloud.google.com/en/feed.atom"
)

// FetchIncidentsJSON retrieves the raw JSON data from the Google Cloud status history page.
func FetchIncidentsJSON() ([]byte, error) {
	return fetchData(jsonURL)
}

// FetchFeedAtom retrieves the raw Atom feed data from the Google Cloud status page.
func FetchFeedAtom() ([]byte, error) {
	return fetchData(atomURL)
}

// fetchData is a helper function to perform the HTTP GET request.
func fetchData(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to get data from %s: %w", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bad status code from %s: %d", url, resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body from %s: %w", url, err)
	}

	return body, nil
}
