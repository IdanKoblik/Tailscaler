package main

import (
	"api/config"
	"api/logger"
	"api/tailscale"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	handler := &tailscale.Handler{}
	router.HandleFunc("/tailscale/create_user", handler.CreateUser).Methods(http.MethodPost)
	router.HandleFunc("/tailscale/get_users", handler.GetAllUsers).Methods(http.MethodGet)
	router.HandleFunc("/tailscale/find_user_by_name/{HostName}", handler.FindByHostName).Methods(http.MethodGet)

	addr, err := config.GetApiURL()
	if err != nil {
		log.Fatalf("Error getting api url: %v\n", err)
		return
	}

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
