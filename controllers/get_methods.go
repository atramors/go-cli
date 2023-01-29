package controllers

import (
	"log"

	"github.com/atramors/go-cli/models"
	"github.com/jmoiron/sqlx"
)

// Get info about provided city.
func GetCityInfo(c string, db *sqlx.DB) models.CityInfo {
	cityObj := models.CityInfo{}
	query := `
		SELECT city_name, temperature, humidity
		FROM "weather"
		WHERE city_name = $1
	`
	err := db.Get(&cityObj, query, c)
	if err != nil {
		log.Println(err)
		cityObj = models.CityInfo{
			CityName:    c,
			Temperature: "Unknown",
			Humidity:    "Unknown",
		}
	}

	log.Printf("Information about weather in %s city\n", c)
	return cityObj
}

// Get info about provided number of cities,
// for example: 10 will print info about 10 cities.
func GetRowsLimitedBy(n string, db *sqlx.DB) []models.CityInfo {
	cityObj := []models.CityInfo{}
	query := `
		SELECT city_name, temperature, humidity
		FROM "weather"
		LIMIT $1
	`
	db.Select(&cityObj, query, n)

	log.Printf("Information about weather in %s cities:\n", n)
	return cityObj
}
