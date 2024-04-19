package tailscale

import (
	"api/utils"
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"strings"
)

type Handler struct {
	Nodes []Node
}

func (h *Handler) CreateNode(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is supported", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	var newNode Node
	err = json.Unmarshal(body, &newNode)
	if err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	// Filter out IPs that do not have a netmask of "/32" for IPv4 or "/128" for IPv6
	filteredIPs := make([]string, 0, len(newNode.AllowedIPs))
	for _, ip := range newNode.AllowedIPs {
		if utils.IsValidIPv4(ip) && strings.HasSuffix(ip, "/32") {
			filteredIPs = append(filteredIPs, ip)
		} else if utils.IsValidIPv6(ip) && strings.HasSuffix(ip, "/128") {
			filteredIPs = append(filteredIPs, ip)
		}
	}

	newNode.AllowedIPs = filteredIPs

	// Remove the old data if a node with the same hostname and router exists
	for i, existingNode := range h.Nodes {
		if existingNode.HostName == newNode.HostName && existingNode.Router == newNode.Router {
			h.Nodes = append(h.Nodes[:i], h.Nodes[i+1:]...)
			break
		}
	}

	h.Nodes = append(h.Nodes, newNode)

	log.Printf("Received Node data: %+v\n", newNode)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write([]byte("Data received successfully"))
	if err != nil {
		http.Error(w, "Failed receiving data:"+err.Error(), http.StatusBadRequest)
		return
	}
}

func (h *Handler) GetAllNodes(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET method is supported", http.StatusMethodNotAllowed)
		return
	}

	if len(h.Nodes) == 0 {
		http.Error(w, "Nodes not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(h.Nodes)
	if err != nil {
		http.Error(w, "Failed to encode nodes data: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) FindNodeByHostName(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET method is supported", http.StatusMethodNotAllowed)
		return
	}

	params := mux.Vars(r)
	var matchedNodes []Node

	for _, node := range h.Nodes {
		if node.HostName == params["HostName"] {
			matchedNodes = append(matchedNodes, node)
		}
	}

	if len(matchedNodes) == 0 {
		http.Error(w, "Node not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(matchedNodes)
	if err != nil {
		http.Error(w, "Failed to encode node data: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
