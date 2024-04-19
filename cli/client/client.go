package client

import (
	"io"
	"log"
	"net/http"
)

type Node struct {
	Router     string   `json:"Router"`
	ID         string   `json:"ID"`
	HostName   string   `json:"HostName"`
	OS         string   `json:"OS"`
	AllowedIPs []string `json:"AllowedIPs"`
	CurAddr    string   `json:"CurAddr"`
	Active     string   `json:"Active"`
}

func createRequest(apiURL string) ([]byte, error) {
	resp, err := http.Get(apiURL)
	if err != nil {
		log.Fatalf("Failed to make request to %s: %s\n", apiURL, err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %s\n", err)
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("API request failed with status code: %d\n", resp.StatusCode)
		return nil, err
	}

	return body, nil
}
