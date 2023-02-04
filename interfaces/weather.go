package interfaces

import (
	"github.com/atramors/go-cli/models"
)

type WeatherCityInfo interface {
	GetByCity(name string) []models.WeatherInfo
	GetByNumber(num string) []models.WeatherInfo
}
