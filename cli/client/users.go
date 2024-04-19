package client

import (
	"encoding/json"
	"fmt"
	"log"
	"tailscaler/config"
)

func GetNodes() ([]*Node, error) {
	nodes, err := getNodes()
	if err != nil {
		log.Fatalf("Error getting nodes: %v\n", err)
		return nil, err
	}

	return nodes, nil
}

func getNodes() ([]*Node, error) {
	url, err := config.GetApiURL()
	if err != nil {
		log.Fatalf("Error getting API url: %v\n", err)
		return nil, err
	}

	apiURL := url + "/tailscale/get_nodes"

	body, err := createRequest(apiURL)
	if err != nil {
		log.Fatalf("Failed to get request: %v\n", err)
		return nil, err
	}

	var nodes []*Node
	err = json.Unmarshal(body, &nodes)
	if err != nil {
		fmt.Printf("Failed to unmarshal JSON response: %s\n", err)
		return nil, err
	}

	return nodes, nil
}
