package main

import (
	"github.com/atramors/go-cli/database"
	"github.com/atramors/go-cli/routes"
)

func main() {
	postgresClient := database.ConnectPostgreSQL()
	DBConn := database.ConnectToDB(postgresClient.DB)
	defer DBConn.CloseConnection()
	routes.StartGin(DBConn)
}
