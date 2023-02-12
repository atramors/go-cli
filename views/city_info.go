package views

import (
	"fmt"

	"github.com/atramors/go-cli/models"
)

func PrintCity(cities []models.WeatherInfo) {
	for _, cityInfo := range cities {
		fmt.Printf(
			"City: %s, temperature: %s, humidity: %s%%\n",
			cityInfo.CityName, cityInfo.Temperature, cityInfo.Humidity)
	}
}
