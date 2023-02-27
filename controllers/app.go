package controllers

// import (
// 	"github.com/atramors/go-cli/database"
// 	"github.com/atramors/go-cli/helpers"
// 	"github.com/atramors/go-cli/models"
// 	"github.com/atramors/go-cli/views"
// )

// type Getter interface {
// 	GetInfo(arg string, query string) (mod []models.WeatherInfo)
// }

// func GetWeather(db Getter, input string, query string) []models.WeatherInfo {
// 	return db.GetInfo(input, query)
// }

// // Main method for getting data from database.
// func GetData(db Getter) {

// 	input := helpers.ArgProcessing()
// 	if input != "" {
// 		info := GetCityInfo(input, db)
// 		views.PrintCity(info)
// 	}
// }

// func GetCityInfo(input string, db Getter) []models.WeatherInfo {
// 	if helpers.ArgIsNum(input) {
// 		return GetWeather(db, input, database.QueryByNumber)
// 	}
// 	return GetWeather(db, input, database.QueryByName)

// }
