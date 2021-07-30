package api

import (
	"OLYMPIC_MEDAL_TABLE/models"
	"encoding/json"
	"net/http"
)

// CountryMedals returns a list of Medals by Country
func CountryMedals(w http.ResponseWriter, r *http.Request) {
	countries := models.Countries

	w.Header().Set("Content-Type", "application/json")
	j, _ := json.Marshal(countries)
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}
