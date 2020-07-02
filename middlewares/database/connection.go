package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func ConnectDatabase() *sql.DB {
	database, err := sql.Open("postgres",
		"postgresql://root@localhost:26257/spyserver?sslmode=disable")
	if err != nil {
		log.Println("Database: No se puede conectar a la base de datos.")
	}

	return database
}
