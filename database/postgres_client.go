package database

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
)

type PostgreSQL struct {
	DB *sqlx.DB
}

// Creates connection to a PostgreSQL Database.
func ConnectPostgreSQL() *PostgreSQL {
	driver := "postgres"
	db_host := os.Getenv(host)
	db_user := os.Getenv(user)
	db_password := os.Getenv(password)
	db_name := os.Getenv(name)
	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		db_host, port, db_user, db_password, db_name,
	)

	db, err := sqlx.Connect(driver, connStr)
	if err != nil {
		log.Println("Failed connect to database.", err)
	}
	return &PostgreSQL{DB: db}
}
