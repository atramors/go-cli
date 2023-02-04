package main

import (
	"github.com/atramors/go-cli/controllers"
	"github.com/atramors/go-cli/database"
)

func main() {
	postgres_client := database.ConnectPostgreSQL()
	db_conn := database.ConnectToDB(postgres_client.DB)
	defer db_conn.CloseConnection()
	controllers.GetData(db_conn)
}
