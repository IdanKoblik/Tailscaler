package utils

import (
	"fmt"
	"strings"
)

type TableNode struct {
	Connections []string
	ID          string   `json:"ID"`
	HostName    string   `json:"HostName"`
	OS          string   `json:"OS"`
	AllowedIPs  []string `json:"AllowedIPs"`
	CurAddr     string   `json:"CurAddr"`
	Active      string   `json:"Active"`
}

func calculateMaxWidths(nodes []*TableNode) (int, int, int, int, int, int) {
	var (
		maxRouterWidth     = 7
		maxIDWidth         = 2
		maxHostNameWidth   = 8
		maxOSWidth         = 3
		maxAllowedIPsWidth = 12
		maxCurAddrWidth    = 12
	)

	for _, node := range nodes {
		maxRouterWidth = max(maxRouterWidth, len(strings.Join(node.Connections, ", ")))
		maxIDWidth = max(maxIDWidth, len(node.ID))
		maxHostNameWidth = max(maxHostNameWidth, len(node.HostName))
		maxOSWidth = max(maxOSWidth, len(node.OS))
		ips := len(strings.Join(node.AllowedIPs, ", "))
		maxAllowedIPsWidth = max(maxAllowedIPsWidth, ips)
		maxCurAddrWidth = max(maxCurAddrWidth, len(node.CurAddr))
	}

	return maxRouterWidth, maxIDWidth, maxHostNameWidth, maxOSWidth, maxAllowedIPsWidth, maxCurAddrWidth
}

func PrintTable(nodes []*TableNode) {
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
		"|\n", "Router", "ID", "NodeName", "OS", "TailscaleIps", "Public IP")

	fmt.Println("+" + strings.Repeat("-", maxRouterWidth+2) + "+" + strings.Repeat("-", maxIDWidth+2) + "+" + strings.Repeat("-", maxHostNameWidth+2) + "+" + strings.Repeat("-", maxOSWidth+2) + "+" + strings.Repeat("-", maxAllowedIPsWidth+2) + "+" + strings.Repeat("-", maxCurAddrWidth+2) + "+--------+")
}

func PrintTableRow(node *TableNode, maxRouterWidth, maxIDWidth, maxHostNameWidth, maxOSWidth, maxAllowedIPsWidth, maxCurAddrWidth int) {
	active := fmt.Sprintf("%v", node.Active)

	fmt.Printf("| %-"+fmt.Sprintf("%ds", maxRouterWidth)+" "+
		"| %-"+fmt.Sprintf("%ds", maxIDWidth)+" "+
		"| %-"+fmt.Sprintf("%ds", maxHostNameWidth)+" "+
		"| %-"+fmt.Sprintf("%ds", maxOSWidth)+" "+
		"| %-"+fmt.Sprintf("%ds", maxAllowedIPsWidth)+" "+
		"| %-"+fmt.Sprintf("%ds", maxCurAddrWidth)+" "+
		"| %-6s "+
		"|\n", strings.Join(node.Connections, ", "), node.ID, node.HostName, node.OS, strings.Join(node.AllowedIPs, ", "), node.CurAddr, active)

	fmt.Println("+" + strings.Repeat("-", maxRouterWidth+2) + "+" + strings.Repeat("-", maxIDWidth+2) + "+" + strings.Repeat("-", maxHostNameWidth+2) + "+" + strings.Repeat("-", maxOSWidth+2) + "+" + strings.Repeat("-", maxAllowedIPsWidth+2) + "+" + strings.Repeat("-", maxCurAddrWidth+2) + "+--------+")
}
