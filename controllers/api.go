package controllers

import (
	"github.com/atramors/go-cli/database"
	"github.com/atramors/go-cli/helpers"
	"github.com/atramors/go-cli/models"
)

type DBGetter interface {
	GetInfo(arg string, query string) (mod []models.WeatherInfo)
}

func GetWeather(db DBGetter, input string, query string) []models.WeatherInfo {
	return db.GetInfo(input, query)
}

// Main method for getting data from database.
func GetData(db DBGetter, input string) []models.WeatherInfo {
	return GetCityInfo(input, db)
}

func GetCityInfo(input string, db DBGetter) []models.WeatherInfo {
	if helpers.ArgIsNum(input) {
		return GetWeather(db, input, database.QueryByNumber)
	}
	return GetWeather(db, input, database.QueryByName)

}
