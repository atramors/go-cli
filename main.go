package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	host     = "DB_HOST"
	port     = 5432
	user     = "GO_DB_USER"
	password = "GO_DB_PASS"
	name     = "GO_DB_NAME"
	driver   = "postgres"
)

var (
	db_host     = os.Getenv(host)
	db_user     = os.Getenv(user)
	db_password = os.Getenv(password)
	db_name     = os.Getenv(name)
	connStr     = fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		db_host, port, db_user, db_password, db_name,
	)
)

type CityInfo struct {
	CityName    string `db:"city_name"`
	Temperature string `db:"temperature"`
	Humidity    string `db:"humidity"`
}

func main() {
	db, err := sqlx.Connect(driver, connStr)
	ErrorCheck(err)
	defer db.Close()

	// first argument - programm name
	if len(os.Args) < 2 {
		fmt.Println("Don't forget to provide a number of rows or city name.")
	} else if len(os.Args) > 2 {
		fmt.Println("For now we support only 1 argument.")
	} else {
		arg := os.Args[1]

		switch {
		case argIsNum(arg):
			fmt.Printf("Information about weather in %s cities:\n", arg)
			response := getRowsLimitedBy(arg, db)
			for _, row := range response {
				fmt.Printf(
					"City: %s, temperature: %s, humidity: %s%%\n",
					row.CityName, row.Temperature, row.Humidity)
			}
		default:
			response := getCityInfo(arg, db)
			fmt.Printf(
				"City: %s, temperature: %s, humidity: %s%%\n",
				response.CityName, response.Temperature, response.Humidity)
		}
	}
}

// Panics if some error occured.
func ErrorCheck(err error) {
	if err != nil {
		panic(err)
	}
}

// Checking if provided string a number.
func argIsNum(s string) bool {
	if len(s) == 0 {
		return false
	}
	digits := "123456789"
	for _, char := range s {
		if !strings.ContainsAny(string(char), digits) {
			return false
		}
	}
	return true
}

// Get info about provided city.
func getCityInfo(c string, db *sqlx.DB) CityInfo {
	cityObj := CityInfo{}
	query := `
		SELECT city_name, temperature, humidity
		FROM "weather"
		WHERE city_name = $1
	`
	err := db.Get(&cityObj, query, c)
	ErrorCheck(err)
	return cityObj

}

// Get info about provided number of cities,
// for example: 10 will print info about 10 cities.
func getRowsLimitedBy(n string, db *sqlx.DB) []CityInfo {
	cityObj := []CityInfo{}
	query := `
		SELECT city_name, temperature, humidity
		FROM "weather"
		LIMIT $1
	`
	db.Select(&cityObj, query, n)
	return cityObj
}
