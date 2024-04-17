package utils

import "tailscaler/client"

func GetNodeConnections(nodes []*client.Node) []string {
	connectionsMap := make(map[string][]string)
	var connections []string

	for _, node := range nodes {
		if _, ok := connectionsMap[node.HostName]; !ok {
			connectionsMap[node.HostName] = []string{}
		}

		connections = append(connections, node.Router)
	}

	return connections

}
