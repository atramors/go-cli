package database

import (
	"log"

	"github.com/atramors/go-cli/models"
	"github.com/jmoiron/sqlx"
)

type Storage struct {
	db *sqlx.DB
}

func ConnectToDB(db *sqlx.DB) *Storage {
	log.Println("Successfull connect to database.")
	return &Storage{db}
}

func (conn *Storage) GetInfo(arg string, query string) (mod []models.WeatherInfo) {
	conn.db.Select(&mod, query, arg)
	return
}

func (conn *Storage) CloseConnection() {
	log.Println("Connection is closed.")
	conn.db.Close()
}
