package helpers

import (
	"log"
	"os"
	"strings"
)

// Processing provided arguments.
func ArgProcessing() (arg string) {
	if len(os.Args) < 2 {
		log.Println("Don't forget to provide a number of rows (more than 0) or city name.")
	} else if len(os.Args) > 2 {
		log.Println("For now we support only 1 argument.")
	} else if os.Args[1] == "0" {
		log.Println("Must be more than 0.")
	} else {
		arg = os.Args[1]
	}
	return
}

// Checking if provided string a number.
func ArgIsNum(s string) bool {
	digits := "1234567890"
	for _, char := range s {
		if !strings.ContainsAny(string(char), digits) {
			return false
		}
	}
	return true
}
