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
	Users []Node
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
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

	var newUser Node
	err = json.Unmarshal(body, &newUser)
	if err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	// Filter out IPs that do not have a netmask of "/32" for IPv4 or "/128" for IPv6
	filteredIPs := make([]string, 0, len(newUser.AllowedIPs))
	for _, ip := range newUser.AllowedIPs {
		if utils.IsValidIPv4(ip) && strings.HasSuffix(ip, "/32") {
			filteredIPs = append(filteredIPs, ip)
		} else if utils.IsValidIPv6(ip) && strings.HasSuffix(ip, "/128") {
			filteredIPs = append(filteredIPs, ip)
		}
	}

	newUser.AllowedIPs = filteredIPs

	// Remove the old data if a user with the same hostname and router exists
	for i, existingUser := range h.Users {
		if existingUser.HostName == newUser.HostName && existingUser.Router == newUser.Router {
			h.Users = append(h.Users[:i], h.Users[i+1:]...)
			break
		}
	}

	h.Users = append(h.Users, newUser)

	log.Printf("Received Node data: %+v\n", newUser)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write([]byte("Data received successfully"))
	if err != nil {
		http.Error(w, "Failed receiving data:"+err.Error(), http.StatusBadRequest)
		return
	}
}

func (h *Handler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET method is supported", http.StatusMethodNotAllowed)
		return
	}

	if len(h.Users) == 0 {
		http.Error(w, "Users not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(h.Users)
	if err != nil {
		http.Error(w, "Failed to encode users data: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) FindByHostName(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET method is supported", http.StatusMethodNotAllowed)
		return
	}

	params := mux.Vars(r)
	var matchedUsers []Node

	for _, user := range h.Users {
		if user.HostName == params["HostName"] {
			matchedUsers = append(matchedUsers, user)
		}
	}

	if len(matchedUsers) == 0 {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(matchedUsers)
	if err != nil {
		http.Error(w, "Failed to encode user data: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
