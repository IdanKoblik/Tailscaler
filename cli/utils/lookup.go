package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"tailscaler/client"
	"tailscaler/config"
)

func LookupUser(hostName string) {
	nodes, err := getUser(hostName)
	if err != nil {
		log.Fatalf("Error getting user: %s :: %v\n", hostName, err)
		return
	}

	PrintTable(nodes)
}

func getUser(hostName string) ([]*client.Node, error) {
	url, err := config.GetApiURL()
	if err != nil {
		log.Fatalf("Error getting API url: %v\n", err)
		return nil, err
	}

	apiURL := fmt.Sprintf("%s/tailscale/find_user_by_name/%s", url, hostName)

	body, err := client.CreateRequest(apiURL)
	if err != nil {
		log.Fatalf("Failed to get request: %v\n", err)
		return nil, err
	}

	var nodes []*client.Node
	err = json.Unmarshal(body, &nodes)
	if err != nil {
		fmt.Printf("Failed to unmarshal JSON response: %s\n", err)
		return nil, err
	}

	return nodes, nil
}
