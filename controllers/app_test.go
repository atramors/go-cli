package controllers

import (
	"reflect"
	"testing"

	"github.com/atramors/go-cli/helpers"
	"github.com/atramors/go-cli/models"
)

type MockStorage struct {
}

func (mt MockStorage) GetInfo(arg string, query string) (mod []models.WeatherInfo) {
	if helpers.ArgIsNum(arg) {
		mod = append(mod,
			models.WeatherInfo{
				CityName:    "Warsaw",
				Temperature: "6",
				Humidity:    "666"},
			models.WeatherInfo{
				CityName:    "Krakow",
				Temperature: "6",
				Humidity:    "666"},
			models.WeatherInfo{
				CityName:    "Rome",
				Temperature: "6",
				Humidity:    "666"},
		)
	} else {
		mod = append(mod, models.WeatherInfo{
			CityName:    arg,
			Temperature: "6",
			Humidity:    "666"})
	}
	return
}

type TestCase struct {
	Name     string
	Input    string
	Expected []models.WeatherInfo
}

func TestGetCityInfo(t *testing.T) {
	testMapper := []TestCase{
		{
			Name:  "If argument is alphabetical",
			Input: "New_York",
			Expected: []models.WeatherInfo{
				{CityName: "New_York", Temperature: "6", Humidity: "666"},
			},
		},
		{
			Name:  "If argument is non alphabetical",
			Input: "3",
			Expected: []models.WeatherInfo{
				{
					CityName:    "Warsaw",
					Temperature: "6",
					Humidity:    "666"},
				{
					CityName:    "Krakow",
					Temperature: "6",
					Humidity:    "666"},
				{
					CityName:    "Rome",
					Temperature: "6",
					Humidity:    "666"},
			},
		},
	}
	for _, test := range testMapper {
		t.Run(test.Name, func(t *testing.T) {
			result := GetCityInfo(test.Input, MockStorage{})
			if !reflect.DeepEqual(result, test.Expected) {
				// if result != test.Expected {
				t.Errorf("'%v' != '%v' with input: '%s'", result, test.Expected, test.Input)
			}
		})
	}
}
