package models

// CountryMedal struct
type CountryMedal struct {
	Position     int64  `json:"position"`
	Country      string `json:"country"`
	GoldMedalds  int64  `json:"goldmedalds"`
	SilverMedals int64  `json:"silvermedalds"`
	BronzeMedals int64  `json:"bronzemedalds"`
}

type CountriesMedals []CountryMedal

var Countries = CountriesMedals{
	{
		Country:      "China",
		GoldMedalds:  17,
		SilverMedals: 20,
		BronzeMedals: 5,
	}, {
		Country:      "USA",
		GoldMedalds:  17,
		SilverMedals: 19,
		BronzeMedals: 19,
	}, {
		Country:      "Japon",
		GoldMedalds:  10,
		SilverMedals: 11,
		BronzeMedals: 9,
	}, {
		Country:      "Colombia",
		GoldMedalds:  17,
		SilverMedals: 10,
		BronzeMedals: 6,
	}, {
		Country:      "Australia",
		GoldMedalds:  10,
		SilverMedals: 11,
		BronzeMedals: 10,
	},
}

func ORGANIZE() {
	for i := 0; i < len(Countries); i++ {
		Countries[i].Position = int64(i + 1)
		for j := 0; j < len(Countries); j++ {
			if Countries[i].GoldMedalds > Countries[j].GoldMedalds {
				auxiliar := Countries[i]
				Countries[i] = Countries[j]
				Countries[j] = auxiliar
			}
			if Countries[i].GoldMedalds == Countries[j].GoldMedalds {
				if Countries[i].SilverMedals > Countries[j].SilverMedals {
					auxiliar := Countries[i]
					Countries[i] = Countries[j]
					Countries[j] = auxiliar
				}
				if Countries[i].SilverMedals == Countries[j].SilverMedals {
					if Countries[i].BronzeMedals > Countries[j].BronzeMedals {
						auxiliar := Countries[i]
						Countries[i] = Countries[j]
						Countries[j] = auxiliar
					}
				}
			}
		}
	}
	for i := 0; i < len(Countries); i++ {
		Countries[i].Position = int64(i + 1)
	}
}
