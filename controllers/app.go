package controllers

import (
	"github.com/atramors/go-cli/database"
	"github.com/atramors/go-cli/helpers"
	"github.com/atramors/go-cli/interfaces"
	"github.com/atramors/go-cli/models"
	"github.com/atramors/go-cli/views"
)

// Main method for getting data from database.
func GetData(db *database.Storage) {

	input := helpers.ArgProcessing()
	if input != "" {
		info := GetCityInfo(input, db)
		views.PrintCity(info)
	}
}

func GetCityInfo(s string, db *database.Storage) []models.WeatherInfo {
	if helpers.ArgIsNum(s) {
		return WeatherInfoForSomeCities(db, s)
	}
	return WeatherInfoByCity(db, s)

}

// Get weather info about provided city (when input - word).
func WeatherInfoByCity(c interfaces.WeatherCityInfo, input string) []models.WeatherInfo {
	return c.GetByCity(input)
}

// Get weather info about provided number of cities (when input - numeric value).
func WeatherInfoForSomeCities(c interfaces.WeatherCityInfo, input string) []models.WeatherInfo {
	return c.GetByNumber(input)
}
