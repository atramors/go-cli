package controllers

import (
	"github.com/atramors/go-cli/database"
	"github.com/atramors/go-cli/helpers"
	"github.com/atramors/go-cli/models"
	"github.com/atramors/go-cli/views"
)

// Main method for getting data from database.
func GetData(db database.DB) {

	input := helpers.ArgProcessing()
	if input != "" {
		info := GetCityInfo(input, db)
		views.PrintCity(info)
	}
}

func GetCityInfo(input string, db database.DB) []models.WeatherInfo {
	if helpers.ArgIsNum(input) {
		return database.GetWeather(db, input, database.QueryByNumber)
	}
	return database.GetWeather(db, input, database.QueryByName)

}
