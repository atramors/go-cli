package models

type WeatherInfo struct {
	CityName    string `db:"city_name"`
	Temperature string `db:"temperature"`
	Humidity    string `db:"humidity"`
}
