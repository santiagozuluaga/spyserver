package database

import (
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func InsertServer(domain string, address string, ssl_grade string, country string, owner string) error {

	database := ConnectDatabase()
	defer database.Close()
	var flag bool

	queryServer :=
		`INSERT INTO
		server (domain, address, ssl_grade, country, owner)
		VALUES ($1, $2, $3, $4, $5) RETURNING true;`

	err := database.QueryRow(queryServer,
		domain, address, ssl_grade, country, owner).Scan(&flag)

	if err != nil {

		fmt.Println(err)
		log.Println("Database: No se pudo ingresar el servidor")
		return err
	}

	return nil
}

func GetServers(host string) ([]Server, error) {

	database := ConnectDatabase()
	defer database.Close()

	var server Server
	var servers []Server

	queryServers :=
		`SELECT address, ssl_grade, country, owner
		FROM server
		WHERE domain = $1;`

	rows, err := database.Query(queryServers, host)
	defer rows.Close()

	if err != nil {

		fmt.Println(err)
		log.Println("Database: No se pudo obtener los datos de los servidores.")
		return servers, err

	}

	for rows.Next() {

		err := rows.Scan(
			&server.Address,
			&server.Ssl_grade,
			&server.Country,
			&server.Owner,
		)

		if err != nil {

			fmt.Println(err)
			log.Println("Database: Error al escanear los datos de los servidores.")
			return servers, err
		}

		servers = append(servers, server)
	}

	return servers, nil
}

func UpdateServer(address string, ssl_grade string) error {

	database := ConnectDatabase()
	defer database.Close()
	var flag bool

	queryServer :=
		`UPDATE server 
		SET ssl_grade = $1
		WHERE address = $2 RETURNING true;;`

	err := database.QueryRow(queryServer,
		ssl_grade, address).Scan(&flag)

	if err != nil {

		fmt.Println(err)
		log.Println("Database: No se pudo actualizar el servidor")
		return err
	}

	return nil

}

func DeleteServer(address string) error {

	database := ConnectDatabase()
	defer database.Close()
	var flag bool

	queryServer :=
		`DELETE
		FROM server
		WHERE address = $1 RETURNING true;`

	err := database.QueryRow(queryServer,
		address).Scan(&flag)

	if err != nil {

		fmt.Println(err)
		log.Println("Database: No se pudo eliminar el servidor")
		return err
	}

	return nil
}
