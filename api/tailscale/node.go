package tailscale

type Node struct {
	Router     string   `json:"Router"`
	ID         string   `json:"ID"`
	HostName   string   `json:"HostName"`
	OS         string   `json:"OS"`
	AllowedIPs []string `json:"AllowedIPs"`
	CurAddr    string   `json:"CurAddr"`
	Active     bool     `json:"Active"`
}
