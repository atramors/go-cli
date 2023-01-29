package views

import (
	"fmt"

	"github.com/atramors/go-cli/models"
)

func PrintCityInfo(c models.CityInfo) {
	fmt.Printf(
		"City: %s, temperature: %s, humidity: %s%%\n",
		c.CityName, c.Temperature, c.Humidity)
}
