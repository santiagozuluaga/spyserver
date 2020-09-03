package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func ConnectDatabase() *sql.DB {
	database, err := sql.Open("postgres",
		"postgresql://root@localhost:26257/spyserver?sslmode=disable")
	if err != nil {
		fmt.Printf("Failed to connect to database: %s\n", err)
	}

	return database
}

func InitDatabase() error {

	querys := []string{
		`CREATE TABLE IF NOT EXISTS domain (
			domain STRING,
			servers_changed BOOL,
			ssl_grade STRING,
			previous_ssl_grade STRING,
			logo STRING,
			title STRING,
			is_down BOOL,
			created STRING,
			updated STRING,
			PRIMARY KEY ("domain")
		);`,
		`CREATE TABLE IF NOT EXISTS server (
			address STRING,
			domain STRING NOT NULL REFERENCES domain(domain) ON DELETE CASCADE,
			ssl_grade STRING,
			country STRING,
			owner STRING,
			PRIMARY KEY (address, domain)
		);`,
	}

	database := ConnectDatabase()
	defer database.Close()

	for i := 0; i < len(querys); i++ {

		_, err := database.Exec(querys[i])
		if err != nil {

			fmt.Println(err)
			log.Println("Database: No se pudo crear las tablas")

			return err
		}
	}

	return nil
}
