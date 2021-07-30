package main

import (
	"OLYMPIC_MEDAL_TABLE/actions"
	"OLYMPIC_MEDAL_TABLE/api"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter().StrictSlash(true)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	server := &http.Server{
		Addr:           ":2000",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// Routes
	r.HandleFunc("/", actions.Home).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/countries", api.CountryMedals).Methods("GET")
	http.Handle("/", r)

	log.Println("Listenig in port 2000...")
	log.Fatal(server.ListenAndServe())
}
