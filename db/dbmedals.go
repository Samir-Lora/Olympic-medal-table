package dbmedals

import (
	"OLYMPIC_MEDAL_TABLE/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DBMEDALS() models.CountriesMedals {
	dsn := "host=localhost user=postgres password=gorm dbname=medals port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	country := models.CountryMedal{}
	if !db.Migrator().HasTable(&country) {
		db.Migrator().CreateTable(&country)
	}
	array := models.CountriesMedals{}
	db.Find(&array)
	for i := 0; i < len(array); i++ {
		array[i].Position = int64(i + 1)
		for j := 0; j < len(array); j++ {
			if array[i].GoldMedalds > array[j].GoldMedalds {
				auxiliar := array[i]
				array[i] = array[j]
				array[j] = auxiliar
			}
			if array[i].GoldMedalds == array[j].GoldMedalds {
				if array[i].SilverMedals > array[j].SilverMedals {
					auxiliar := array[i]
					array[i] = array[j]
					array[j] = auxiliar
				}
				if array[i].SilverMedals == array[j].SilverMedals {
					if array[i].BronzeMedals > array[j].BronzeMedals {
						auxiliar := array[i]
						array[i] = array[j]
						array[j] = auxiliar
					}
				}
			}
		}
	}
	for i := 0; i < len(array); i++ {
		array[i].Position = int64(i + 1)
		array[i].AllMedals = array[i].GoldMedalds + array[i].SilverMedals + array[i].BronzeMedals
	}
	return array

}
