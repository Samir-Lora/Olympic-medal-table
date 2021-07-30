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
	organice()
	j, _ := json.Marshal(countries)
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func organice() {
	for i := 0; i < len(models.Countries); i++ {
		for j := 0; j < len(models.Countries); j++ {
			if models.Countries[i].GoldMedalds > models.Countries[j].GoldMedalds {
				auxiliar := models.Countries[i]
				models.Countries[i] = models.Countries[j]
				models.Countries[j] = auxiliar

			}
			if models.Countries[i].GoldMedalds == models.Countries[j].GoldMedalds {
				if models.Countries[i].SilverMedals > models.Countries[j].SilverMedals {
					auxiliar := models.Countries[i]
					models.Countries[i] = models.Countries[j]
					models.Countries[j] = auxiliar
				}
				if models.Countries[i].SilverMedals == models.Countries[j].SilverMedals {
					if models.Countries[i].BronzeMedals > models.Countries[j].BronzeMedals {
						auxiliar := models.Countries[i]
						models.Countries[i] = models.Countries[j]
						models.Countries[j] = auxiliar
					}
				}
			}
		}
	}
}

// Debemos definir la variable que alojará los valores decodificados
