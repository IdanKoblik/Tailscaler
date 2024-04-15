package tailscale

import (
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

	var user Node
	err = json.Unmarshal(body, &user)
	if err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	existingIPs := make(map[string]bool)

	for _, existingUser := range h.Users {
		for _, existingIP := range existingUser.AllowedIPs {
			existingIPs[existingIP] = true
		}
	}

	for _, ip := range user.AllowedIPs {
		if existingIPs[ip] {
			http.Error(w, "User already exists", http.StatusConflict)
			return
		}

		// Filter out IPs that do not have a netmask of "/32"
		filteredIPs := make([]string, 0, len(user.AllowedIPs))
		for _, ip := range user.AllowedIPs {
			if strings.HasSuffix(ip, "/32") {
				filteredIPs = append(filteredIPs, ip)
				user.AllowedIPs = filteredIPs
			}
		}
	}

	h.Users = append(h.Users, user)

	log.Printf("Received Node data: %+v\n", user)

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
	for _, user := range h.Users {
		if user.HostName == params["HostName"] {
			w.Header().Set("Content-Type", "application/json")
			err := json.NewEncoder(w).Encode(user)
			if err != nil {
				http.Error(w, "Failed to encode user data: "+err.Error(), http.StatusInternalServerError)
				return
			}
			return
		}
	}

	http.Error(w, "User not found", http.StatusNotFound)
}
