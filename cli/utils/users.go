package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"slices"
	"tailscaler/client"
	"tailscaler/config"
)

func GetUsers() {
	nodes, err := getUsers()
	if err != nil {
		log.Fatalf("Error getting users: %v\n", err)
		return
	}

	PrintTable(nodes)
}

func getUsers() ([]*TableNode, error) {
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

	var nodes []*client.Node
	err = json.Unmarshal(body, &nodes)
	if err != nil {
		fmt.Printf("Failed to unmarshal JSON response: %s\n", err)
		return nil, err
	}

	var tableNodes []*TableNode
	for _, node := range nodes {
		found := false
		for _, existingNode := range tableNodes {
			if existingNode.ID == node.ID {
				found = true

				// Update if not already present
				for _, ip := range node.AllowedIPs {
					if !slices.Contains(existingNode.AllowedIPs, ip) {
						existingNode.AllowedIPs = append(existingNode.AllowedIPs, ip)
					}
				}

				existingNode.CurAddr = node.CurAddr
				existingNode.Active = node.Active
				break
			}
		}

		if !found {
			tableNode := &TableNode{
				Connections: GetNodeConnections(nodes),
				ID:          node.ID,
				HostName:    node.HostName,
				OS:          node.OS,
				AllowedIPs:  node.AllowedIPs,
				CurAddr:     node.CurAddr,
				Active:      node.Active,
			}
			tableNodes = append(tableNodes, tableNode)
		}
	}

	return tableNodes, nil
}
