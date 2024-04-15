package main

import (
	"api/logger"
	"api/tailscale"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

type ServerConfig struct {
	Ip   string `json:"ip"`
	Port string `json:"port"`
}

func main() {
	file, err := os.ReadFile("config.json")
	if err != nil {
		log.Fatal("Error reading config file: ", err)
		return
	}

	var config ServerConfig
	err = json.Unmarshal(file, &config)
	if err != nil {
		log.Fatal("Error parsing JSON: ", err)
		return
	}

	router := mux.NewRouter()

	handler := &tailscale.Handler{}
	router.HandleFunc("/tailscale/create_user", handler.CreateUser).Methods(http.MethodPost)
	router.HandleFunc("/tailscale/get_users", handler.GetAllUsers).Methods(http.MethodGet)
	router.HandleFunc("/tailscale/find_user_by_name/{HostName}", handler.FindByHostName).Methods(http.MethodGet)

	addr := config.Ip + ":" + config.Port
	server := http.Server{
		Addr:    addr,
		Handler: logger.Log(router),
	}

	log.Printf("Starting server. %s:\n", addr)
	err = server.ListenAndServe()
	if err != nil {
		log.Fatalf("Error acure while trying to listen and server to the server: %v\n", err)
		return
	}
}
