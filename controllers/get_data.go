package controllers

import (
	"log"
	"os"

	"github.com/atramors/go-cli/helpers"
	"github.com/atramors/go-cli/initializers"
	"github.com/atramors/go-cli/views"
)

// Main method for getting data from database.
// Using these methods:
//
// GetCityInfo -> city name have been provided (Example: "Barcelona").
//
// GetRowsLimitedBy -> number of cities have been provided (Example: 8).
func GetData() {
	// first argument - programm name
	if len(os.Args) < 2 {
		log.Println("Don't forget to provide a number of rows or city name.")
	} else if len(os.Args) > 2 {
		log.Println("For now we support only 1 argument.")
	} else {
		arg := os.Args[1]
		switch {
		case arg == "0":
			log.Println("Must be more than 0.")
		case helpers.ArgIsNum(arg):
			response := GetRowsLimitedBy(arg, initializers.DB_CONN)
			for _, row := range response {
				views.PrintCityInfo(row)
			}
		case !helpers.ArgIsNum(arg):
			response := GetCityInfo(arg, initializers.DB_CONN)
			views.PrintCityInfo(response)
		default:
			log.Printf("For now we didn't support this option: %s\n", arg)
		}
	}
}
