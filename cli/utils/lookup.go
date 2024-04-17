package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"slices"
	"tailscaler/client"
	"tailscaler/config"
)

func LookupUser(hostName string) {
	nodes, err := getUser(hostName)
	if err != nil {
		log.Fatalf("Error getting user: %s :: %v\n", hostName, err)
		return
	}

	var tableNodes []*TableNode
	tableNodes = append(tableNodes, &nodes)

	PrintTable(tableNodes)
}

func getUser(hostName string) (TableNode, error) {
	url, err := config.GetApiURL()
	if err != nil {
		log.Fatalf("Error getting API url: %v\n", err)
		return TableNode{}, err
	}

	apiURL := fmt.Sprintf("%s/tailscale/find_user_by_name/%s", url, hostName)

	body, err := client.CreateRequest(apiURL)
	if err != nil {
		log.Fatalf("Failed to get request: %v\n", err)
		return TableNode{}, err
	}

	var nodes []*client.Node
	err = json.Unmarshal(body, &nodes)
	if err != nil {
		fmt.Printf("Failed to unmarshal JSON response: %s\n", err)
		return TableNode{}, err
	}

	var id string
	var hostname string
	var os string
	var curAddr string
	var active string
	var allowedIPs []string

	for _, node := range nodes {
		id = node.ID
		hostname = node.HostName
		os = node.OS

		if len(node.CurAddr) != 0 {
			curAddr = node.CurAddr
		}

		if len(node.Active) != 0 {
			active = node.Active
		}

		for _, ip := range node.AllowedIPs {
			if !slices.Contains(allowedIPs, ip) {
				allowedIPs = append(allowedIPs, ip)
			}
		}
	}

	tableNode := TableNode{
		Connections: GetNodeConnections(nodes),
		ID:          id,
		HostName:    hostname,
		OS:          os,
		AllowedIPs:  allowedIPs,
		CurAddr:     curAddr,
		Active:      active,
	}

	return tableNode, nil
}
