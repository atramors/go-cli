package initializers

import (
	"fmt"
	"log"
	"os"

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
	DB_CONN *sqlx.DB
)

// Connecting to Database
func ConnectToDB() {
	var connectionErr error
	DB_CONN, connectionErr = sqlx.Connect(driver, connStr)
	if connectionErr != nil {
		log.Println("Failed connect to database.")
	}
}
