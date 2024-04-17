package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"tailscaler/client"
	"tailscaler/config"
)

func GetUsers() {
	nodes, err := getUsers()
	if err != nil {
		log.Fatalf("Error getting users: %v\n", err)
		return
	}

	var nodePointers []*client.Node
	for _, node := range nodes {
		nodePointers = append(nodePointers, &node)
	}

	PrintTable(nodePointers)
}

func getUsers() ([]client.Node, error) {
	url, err := config.GetApiURL()
	if err != nil {
		log.Fatalf("Error getting API url: %v\n", err)
		return nil, err
	}

	apiURL := url + "/tailscale/get_users"

	body, err := client.CreateRequest(apiURL)
	if err != nil {
		log.Fatalf("Failed to get request: %v\n", err)
		return nil, err
	}

	var nodes []client.Node
	err = json.Unmarshal(body, &nodes)
	if err != nil {
		fmt.Printf("Failed to unmarshal JSON response: %s\n", err)
		return nil, err
	}

	return nodes, nil
}
