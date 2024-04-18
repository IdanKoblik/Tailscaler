package utils

import (
	"fmt"
	"strings"
	"tailscaler/client"
)

func calculateMaxWidths(nodes []*client.Node) (int, int, int, int, int, int) {
	var (
		maxRouterWidth     = 19
		maxIDWidth         = 2
		maxHostNameWidth   = 8
		maxOSWidth         = 3
		maxAllowedIPsWidth = 12
		maxCurAddrWidth    = 12
	)

	for _, node := range nodes {
		maxRouterWidth = max(maxRouterWidth, len(node.Router))
		maxIDWidth = max(maxIDWidth, len(node.ID))
		maxHostNameWidth = max(maxHostNameWidth, len(node.HostName))
		maxOSWidth = max(maxOSWidth, len(node.OS))
		ips := len(strings.Join(node.AllowedIPs, ", "))
		maxAllowedIPsWidth = max(maxAllowedIPsWidth, ips)
		maxCurAddrWidth = max(maxCurAddrWidth, len(node.CurAddr))
	}

	return maxRouterWidth, maxIDWidth, maxHostNameWidth, maxOSWidth, maxAllowedIPsWidth, maxCurAddrWidth
}

func PrintTable(nodes []*client.Node) {
	if len(nodes) == 0 {
		fmt.Println("No nodes to print.")
		return
	}

	maxRouterWidth, maxIDWidth, maxHostNameWidth, maxOSWidth, maxAllowedIPsWidth, maxCurAddrWidth := calculateMaxWidths(nodes)
	printTableHeader(maxRouterWidth, maxIDWidth, maxHostNameWidth, maxOSWidth, maxAllowedIPsWidth, maxCurAddrWidth)

	for _, node := range nodes {
		PrintTableRow(node, maxRouterWidth, maxIDWidth, maxHostNameWidth, maxOSWidth, maxAllowedIPsWidth, maxCurAddrWidth)
	}
}

func printTableHeader(maxRouterWidth, maxIDWidth, maxHostNameWidth, maxOSWidth, maxAllowedIPsWidth, maxCurAddrWidth int) {
	fmt.Println("+" + strings.Repeat("-", maxRouterWidth+2) + "+" + strings.Repeat("-", maxIDWidth+2) + "+" + strings.Repeat("-", maxHostNameWidth+2) + "+" + strings.Repeat("-", maxOSWidth+2) + "+" + strings.Repeat("-", maxAllowedIPsWidth+2) + "+" + strings.Repeat("-", maxCurAddrWidth+2) + "+--------+")

	fmt.Printf("| %-"+fmt.Sprintf("%ds", maxRouterWidth)+" "+
		"| %-"+fmt.Sprintf("%ds", maxIDWidth)+" "+
		"| %-"+fmt.Sprintf("%ds", maxHostNameWidth)+" "+
		"| %-"+fmt.Sprintf("%ds", maxOSWidth)+" "+
		"| %-"+fmt.Sprintf("%ds", maxAllowedIPsWidth)+" "+
		"| %-"+fmt.Sprintf("%ds", maxCurAddrWidth)+" "+
		"| Active "+
		"|\n", "Connected routeres", "ID", "NodeName", "OS", "TailscaleIps", "Public IP")

	fmt.Println("+" + strings.Repeat("-", maxRouterWidth+2) + "+" + strings.Repeat("-", maxIDWidth+2) + "+" + strings.Repeat("-", maxHostNameWidth+2) + "+" + strings.Repeat("-", maxOSWidth+2) + "+" + strings.Repeat("-", maxAllowedIPsWidth+2) + "+" + strings.Repeat("-", maxCurAddrWidth+2) + "+--------+")
}

func PrintTableRow(node *client.Node, maxRouterWidth, maxIDWidth, maxHostNameWidth, maxOSWidth, maxAllowedIPsWidth, maxCurAddrWidth int) {
	active := fmt.Sprintf("%v", node.Active)

	fmt.Printf("| %-"+fmt.Sprintf("%ds", maxRouterWidth)+" "+
		"| %-"+fmt.Sprintf("%ds", maxIDWidth)+" "+
		"| %-"+fmt.Sprintf("%ds", maxHostNameWidth)+" "+
		"| %-"+fmt.Sprintf("%ds", maxOSWidth)+" "+
		"| %-"+fmt.Sprintf("%ds", maxAllowedIPsWidth)+" "+
		"| %-"+fmt.Sprintf("%ds", maxCurAddrWidth)+" "+
		"| %-6s "+
		"|\n", node.Router, node.ID, node.HostName, node.OS, strings.Join(node.AllowedIPs, ", "), node.CurAddr, active)

	fmt.Println("+" + strings.Repeat("-", maxRouterWidth+2) + "+" + strings.Repeat("-", maxIDWidth+2) + "+" + strings.Repeat("-", maxHostNameWidth+2) + "+" + strings.Repeat("-", maxOSWidth+2) + "+" + strings.Repeat("-", maxAllowedIPsWidth+2) + "+" + strings.Repeat("-", maxCurAddrWidth+2) + "+--------+")
}
