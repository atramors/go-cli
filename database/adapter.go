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

func (conn *Storage) CloseConnection() {
	log.Println("Connection is closed.")
	conn.db.Close()
}

func (conn *Storage) GetByCity(name string) (mod []models.WeatherInfo) {
	conn.db.Select(&mod, QueryByName, name)
	return
}

func (conn *Storage) GetByNumber(num string) (mod []models.WeatherInfo) {
	conn.db.Select(&mod, QueryByNumber, num)
	return
}
