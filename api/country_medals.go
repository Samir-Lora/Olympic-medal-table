package api

import (
	dbmedals "OLYMPIC_MEDAL_TABLE/db"

	"encoding/json"
	"net/http"
)

// CountryMedals returns a list of Medals by Country
//func CountryMedals(w http.ResponseWriter, r *http.Request) {
//countries := models.Countries
//models.ORGANIZE()
//w.Header().Set("Content-Type", "application/json")
//j, _ := json.Marshal(countries)
//w.WriteHeader(http.StatusOK)
//w.Write(j)
//}

func CountryMedals(w http.ResponseWriter, r *http.Request) {
	countrymedals := dbmedals.DBMEDALS()
	w.Header().Set("Content-Type", "application/json")
	j, _ := json.Marshal(countrymedals)
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}
