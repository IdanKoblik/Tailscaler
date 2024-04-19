package utils

import (
	"io"
	"os"
	client "tailscaler/client"
	"testing"
)

func TestCalculateMaxWidths(t *testing.T) {
	node := client.Node{
		Router:     "Router2",
		ID:         "71",
		HostName:   "Test2",
		OS:         "arch",
		AllowedIPs: []string{"10.5.2.29/32"},
		CurAddr:    "1.1.1.1",
		Active:     true,
	}

	nodes := []*client.Node{&node}

	tests := []struct {
		Name       string
		router     int
		ID         int
		HostName   int
		OS         int
		AllowedIps int
		CurAddr    int
	}{
		{"Expected widths", 9, 2, 8, 4, 12, 12},
	}

	for _, test := range tests {
		maxRouterWidth, maxIDWidth, maxHostNameWidth, maxOSWidth, maxAllowedIPsWidth, maxCurAddrWidth := calculateMaxWidths(nodes)

		if maxRouterWidth != test.router {
			t.Errorf("expected router width %d, got %d", test.router, maxRouterWidth)
		}
		if maxIDWidth != test.ID {
			t.Errorf("expected ID width %d, got %d", test.ID, maxIDWidth)
		}
		if maxHostNameWidth != test.HostName {
			t.Errorf("expected HostName width %d, got %d", test.HostName, maxHostNameWidth)
		}
		if maxOSWidth != test.OS {
			t.Errorf("expected OS width %d, got %d", test.OS, maxOSWidth)
		}
		if maxAllowedIPsWidth != test.AllowedIps {
			t.Errorf("expected AllowedIPs width %d, got %d", test.AllowedIps, maxAllowedIPsWidth)
		}
		if maxCurAddrWidth != test.CurAddr {
			t.Errorf("expected CurAddr width %d, got %d", test.CurAddr, maxCurAddrWidth)
		}
	}
}

func TestPrintTableHeader(t *testing.T) {
	// Redirect standard output to capture printed output
	old := os.Stdout
	r, w, _ := os.Pipe()

	os.Stdout = w
	node := client.Node{
		Router:     "Router2",
		ID:         "71",
		HostName:   "Test2",
		OS:         "arch",
		AllowedIPs: []string{"10.5.2.29/32"},
		CurAddr:    "1.1.0.1",
		Active:     false,
	}

	nodes := []*client.Node{&node}
	maxRouterWidth, maxIDWidth, maxHostNameWidth, maxOSWidth, maxAllowedIPsWidth, maxCurAddrWidth := calculateMaxWidths(nodes)
	printTableHeader(maxRouterWidth, maxIDWidth, maxHostNameWidth, maxOSWidth, maxAllowedIPsWidth, maxCurAddrWidth)

	// Restore standard output
	w.Close()
	os.Stdout = old

	// Read captured output
	out, _ := io.ReadAll(r)

	expected := "+-----------+-----------+-----------+----+----------+------+--------------+--------------+--------+\n| Router1   | Router2   | Router3   | ID | NodeName | OS   | TailscaleIps | Public IP    | Active |\n+-----------+-----------+-----------+----+----------+------+--------------+--------------+--------+\n"
	if string(out) != expected {
		t.Fatalf("Expected %q\n got %q", expected, out)
	}
}

func TestPrintTableRow(t *testing.T) {
	// Redirect standard output to capture printed output
	old := os.Stdout
	r, w, _ := os.Pipe()

	os.Stdout = w
	node := client.Node{
		Router:     "Router2",
		ID:         "71",
		HostName:   "Test2",
		OS:         "arch",
		AllowedIPs: []string{"10.5.2.29/32"},
		CurAddr:    "1.1.0.1",
		Active:     false,
	}

	nodes := []*client.Node{&node}
	maxRouterWidth, maxIDWidth, maxHostNameWidth, maxOSWidth, maxAllowedIPsWidth, maxCurAddrWidth := calculateMaxWidths(nodes)
	PrintTableRow(&node, maxRouterWidth, maxIDWidth, maxHostNameWidth, maxOSWidth, maxAllowedIPsWidth, maxCurAddrWidth)

	// Restore standard output
	w.Close()
	os.Stdout = old

	// Read captured output
	out, _ := io.ReadAll(r)

	expected := "|           | \x1b[32mConnected\x1b[37m |           | 71 | Test2    | arch | 10.5.2.29/32 | 1.1.0.1      | false  |\n+-----------+-----------+-----------+----+----------+------+--------------+--------------+--------+\n"
	if string(out) != expected {
		t.Fatalf("Expected %q\n got %q", expected, out)
	}
}

func TestPrintTable(t *testing.T) {
	// Redirect standard output to capture printed output
	old := os.Stdout
	r, w, _ := os.Pipe()

	os.Stdout = w
	node := client.Node{
		Router:     "Router2",
		ID:         "71",
		HostName:   "Test2",
		OS:         "arch",
		AllowedIPs: []string{"10.5.2.29/32"},
		CurAddr:    "1.1.0.1",
		Active:     false,
	}

	nodes := []*client.Node{&node}
	PrintTable(nodes)

	// Restore standard output
	w.Close()
	os.Stdout = old

	// Read captured output
	out, _ := io.ReadAll(r)

	expected := "+-----------+-----------+-----------+----+----------+------+--------------+--------------+--------+\n| Router1   | Router2   | Router3   | ID | NodeName | OS   | TailscaleIps | Public IP    | Active |\n+-----------+-----------+-----------+----+----------+------+--------------+--------------+--------+\n|           | \u001B[32mConnected\u001B[37m |           | 71 | Test2    | arch | 10.5.2.29/32 | 1.1.0.1      | false  |\n+-----------+-----------+-----------+----+----------+------+--------------+--------------+--------+\n"
	if string(out) != expected {
		t.Fatalf("Expected %q\n got %q", expected, out)
	}
}
