package controllers

import (
	"github.com/atramors/go-cli/models"
)

type DBGetter interface {
	GetInfo(arg string) (mod []models.WeatherInfo)
}

func GetWeather(db DBGetter, input string) []models.WeatherInfo {
	return db.GetInfo(input)
}

// Main method for getting data from database.
func GetData(db DBGetter, input string) []models.WeatherInfo {
	return GetWeather(db, input)
}
