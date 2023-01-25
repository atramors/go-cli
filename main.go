package main

import (
	"database/sql"
	"fmt"
	"os"
	"unicode"

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

var (
	cityName    string
	temperature string
	humidity    string
)

func main() {
	db, err := sql.Open(driver, connStr)
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
			getRowsLimitedBy(arg, db)
		default:
			getCityInfo(arg, db)
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
	for _, char := range s {
		if !unicode.IsDigit(rune(char)) {
			return false
		}
	}
	return true
}

// Get info about provided city.
func getCityInfo(c string, db *sql.DB) {
	query := fmt.Sprintf(
		`SELECT city_name, temperature, humidity
		 FROM "weather"
		 WHERE city_name = '%s'`, c)
	err := db.QueryRow(query).Scan(&cityName, &temperature, &humidity)
	ErrorCheck(err)

	fmt.Printf(
		"City: %s, temperature: %s, humidity: %s\n",
		cityName, temperature, humidity)
}

// Get info about provided number of cities,
// for example: 10 will print info about 10 cities.
func getRowsLimitedBy(n string, db *sql.DB) {
	query := fmt.Sprintf(
		`SELECT city_name, temperature, humidity
		 FROM "weather"
		 LIMIT '%s'`, n)
	rows, err := db.Query(query)
	ErrorCheck(err)
	fmt.Printf("Information about weather in %s cities:\n", n)
	for rows.Next() {
		rows.Scan(&cityName, &temperature, &humidity)
		// fmt.Println(cityName, temperature, humidity)
		fmt.Printf(
			"City: %s, temperature: %s, humidity: %s\n",
			cityName, temperature, humidity)
	}
}
