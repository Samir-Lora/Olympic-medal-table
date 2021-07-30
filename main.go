package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/index.html")
}

func main() {
	r := mux.NewRouter().StrictSlash(true)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
	server := &http.Server{
		Addr:           ":2000",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	r.HandleFunc("/", Home).Methods("GET", "OPTIONS")
	log.Println("Listenig...")
	log.Fatal(server.ListenAndServe())
}
