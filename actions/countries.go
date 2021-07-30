package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type CountryMedal struct {
	Position     int64  `json:"position"`
	Country      string `json:"country"`
	GoldMedalds  int64  `json:"goldmedalds"`
	SilverMedals int64  `json:"silvermedalds"`
	BronzeMedals int64  `json:"bronzemedalds"`
}

type CountriesMedal []CountryMedal

func main() {
	r := mux.NewRouter().StrictSlash(true)
	server := &http.Server{
		Addr:           ":1000",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	r.HandleFunc("/countries", Getjson).Methods("GET", "OPTIONS")
	log.Println("Listenig...")
	log.Fatal(server.ListenAndServe())
}

func setupCorsResponse(w *http.ResponseWriter, r *http.Request) {

	(*w).Header().Set("Content-Type", "application/json")
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func Getjson(w http.ResponseWriter, r *http.Request) {
	setupCorsResponse(&w, r)
	countries := CountriesMedal{
		{
			Position:     1,
			Country:      "China",
			GoldMedalds:  13,
			SilverMedals: 4,
			BronzeMedals: 5,
		}, {
			Position:     2,
			Country:      "USA",
			GoldMedalds:  12,
			SilverMedals: 6,
			BronzeMedals: 9,
		}, {
			Position:     3,
			Country:      "Japon",
			GoldMedalds:  11,
			SilverMedals: 11,
			BronzeMedals: 9,
		}, {
			Position:     4,
			Country:      "Colombia",
			GoldMedalds:  7,
			SilverMedals: 10,
			BronzeMedals: 6,
		}, {
			Position:     5,
			Country:      "Australia",
			GoldMedalds:  6,
			SilverMedals: 1,
			BronzeMedals: 9,
		},
	}
	w.Header().Set("Content-Type", "application/json")
	j, _ := json.Marshal(countries)
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}
