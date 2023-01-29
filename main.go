package main

import (
	"github.com/atramors/go-cli/controllers"
	"github.com/atramors/go-cli/initializers"
)

func init() {
	initializers.ConnectToDB()
}

func main() {
	defer initializers.DB_CONN.Close()
	controllers.GetData()
}
