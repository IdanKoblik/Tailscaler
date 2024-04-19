package utils

import (
	"fmt"
	"strings"
	"tailscaler/client"
	"tailscaler/math"
)

type Table struct {
	Connection1 string
	Connection2 string
	Connection3 string
	ID          string
	HostName    string
	OS          string
	AllowedIPs  []string
	CurAddr     string
	Active      string
}

func calculateMaxWidths(nodes []*client.Node) (int, int, int, int, int, int) {
	var (
		maxRouterWidth     = 9
		maxIDWidth         = 2
		maxHostNameWidth   = 8
		maxOSWidth         = 3
		maxAllowedIPsWidth = 12
		maxCurAddrWidth    = 12
	)

	for _, node := range nodes {
		maxRouterWidth = math.Max(maxRouterWidth, len(node.Router))
		maxIDWidth = math.Max(maxIDWidth, len(node.ID))
		maxHostNameWidth = math.Max(maxHostNameWidth, len(node.HostName))
		maxOSWidth = math.Max(maxOSWidth, len(node.OS))
		ips := len(strings.Join(node.AllowedIPs, ", "))
		maxAllowedIPsWidth = math.Max(maxAllowedIPsWidth, ips)
		maxCurAddrWidth = math.Max(maxCurAddrWidth, len(node.CurAddr))
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
	fmt.Println("+" + strings.Repeat("-", maxRouterWidth+2) + "+" + strings.Repeat("-", maxRouterWidth+2) + "+" + strings.Repeat("-", maxRouterWidth+2) + "+" + strings.Repeat("-", maxIDWidth+2) + "+" + strings.Repeat("-", maxHostNameWidth+2) + "+" + strings.Repeat("-", maxOSWidth+2) + "+" + strings.Repeat("-", maxAllowedIPsWidth+2) + "+" + strings.Repeat("-", maxCurAddrWidth+2) + "+--------+")

	fmt.Printf("| %-"+fmt.Sprintf("%ds", maxRouterWidth)+" "+
		"| %-"+fmt.Sprintf("%ds", maxRouterWidth)+" "+
		"| %-"+fmt.Sprintf("%ds", maxRouterWidth)+" "+
		"| %-"+fmt.Sprintf("%ds", maxIDWidth)+" "+
		"| %-"+fmt.Sprintf("%ds", maxHostNameWidth)+" "+
		"| %-"+fmt.Sprintf("%ds", maxOSWidth)+" "+
		"| %-"+fmt.Sprintf("%ds", maxAllowedIPsWidth)+" "+
		"| %-"+fmt.Sprintf("%ds", maxCurAddrWidth)+" "+
		"| Active "+
		"|\n", "Router1", "Router2", "Router3", "ID", "NodeName", "OS", "TailscaleIps", "Public IP")

	fmt.Println("+" + strings.Repeat("-", maxRouterWidth+2) + "+" + strings.Repeat("-", maxRouterWidth+2) + "+" + strings.Repeat("-", maxRouterWidth+2) + "+" + strings.Repeat("-", maxIDWidth+2) + "+" + strings.Repeat("-", maxHostNameWidth+2) + "+" + strings.Repeat("-", maxOSWidth+2) + "+" + strings.Repeat("-", maxAllowedIPsWidth+2) + "+" + strings.Repeat("-", maxCurAddrWidth+2) + "+--------+")
}

func PrintTableRow(node *client.Node, maxRouterWidth, maxIDWidth, maxHostNameWidth, maxOSWidth, maxAllowedIPsWidth, maxCurAddrWidth int) {
	active := fmt.Sprintf("%v", node.Active)
	var connection1, connection2, connection3 string
	routerToConnection := map[string]*string{
		"Router1": &connection1,
		"Router2": &connection2,
		"Router3": &connection3,
	}

	if conn, exists := routerToConnection[node.Router]; exists {
		//*conn = node.Router
		// \033[92m - Green color
		*conn = "\033[92mConnected\033[0m"
	}

	table := Table{
		Connection1: connection1,
		Connection2: connection2,
		Connection3: connection3,
		ID:          node.ID,
		HostName:    node.HostName,
		OS:          node.OS,
		AllowedIPs:  node.AllowedIPs,
		CurAddr:     node.CurAddr,
		Active:      active,
	}

	fmt.Printf("| %-"+fmt.Sprintf("%ds", maxRouterWidth)+" "+
		"| %-"+fmt.Sprintf("%ds", maxRouterWidth)+" "+
		"| %-"+fmt.Sprintf("%ds", maxRouterWidth)+" "+
		"| %-"+fmt.Sprintf("%ds", maxIDWidth)+" "+
		"| %-"+fmt.Sprintf("%ds", maxHostNameWidth)+" "+
		"| %-"+fmt.Sprintf("%ds", maxOSWidth)+" "+
		"| %-"+fmt.Sprintf("%ds", maxAllowedIPsWidth)+" "+
		"| %-"+fmt.Sprintf("%ds", maxCurAddrWidth)+" "+
		"| %-6s "+
		"|\n",
		table.Connection1,
		table.Connection2,
		table.Connection3,
		table.ID,
		table.HostName,
		table.OS,
		strings.Join(table.AllowedIPs, ", "),
		table.CurAddr,
		active,
	)

	fmt.Println("+" + strings.Repeat("-", maxRouterWidth+2) + "+" + strings.Repeat("-", maxRouterWidth+2) + "+" + strings.Repeat("-", maxRouterWidth+2) + "+" + strings.Repeat("-", maxIDWidth+2) + "+" + strings.Repeat("-", maxHostNameWidth+2) + "+" + strings.Repeat("-", maxOSWidth+2) + "+" + strings.Repeat("-", maxAllowedIPsWidth+2) + "+" + strings.Repeat("-", maxCurAddrWidth+2) + "+--------+")
}
