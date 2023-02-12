package database

const (
	QueryByName = `
		SELECT city_name, temperature, humidity
		FROM "weather"
		WHERE city_name = $1;
	`
	QueryByNumber = `
		SELECT city_name, temperature, humidity
		FROM "weather"
		LIMIT $1;
	`
)
