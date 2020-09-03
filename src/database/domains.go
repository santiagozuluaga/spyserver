package database

import (
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func InsertDomain(domain string, servers_changed bool, ssl_grade string, previous_ssl_grade string, logo string, title string, is_down bool, updated string) error {

	database := ConnectDatabase()
	defer database.Close()
	var response string

	queryDomain :=
		`INSERT INTO domain 
			(domain, servers_changed, ssl_grade, previous_ssl_grade, logo, title, is_down, updated)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING domain.domain;`

	err := database.QueryRow(queryDomain,
		domain,
		servers_changed,
		ssl_grade,
		previous_ssl_grade,
		logo,
		title,
		is_down,
		updated,
	).Scan(&response)

	if err != nil {

		fmt.Println(err)
		log.Println("Database: No se pudo ingresar el dominio")
		return err
	}

	/*
		for i := 0; i < len(domain.Info.Servers); i++ {

			InsertServer(domain.Domain, domain.Info.Servers[i].Address, domain.Info.Servers[i].Ssl_grade, domain.Info.Servers[i].Country, domain.Info.Servers[i].Owner)
		}
	*/

	return nil

}

func GetDomain(host string) (Domain, error) {

	database := ConnectDatabase()
	defer database.Close()

	var domain Domain

	queryDomain :=
		`SELECT *
		FROM domain 
		WHERE domain = $1;`

	rows := database.QueryRow(queryDomain, host)
	err := rows.Scan(
		&domain.Domain,
		&domain.Info.Servers_changed,
		&domain.Info.Ssl_grade,
		&domain.Info.Previous_ssl_grade,
		&domain.Info.Logo,
		&domain.Info.Title,
		&domain.Info.Is_down,
		&domain.Updated,
	)

	if err != nil {

		fmt.Println(err)
		log.Println("Database: No se pudo obtener el dominio")
		return Domain{}, err
	}

	domain.Info.Servers, err = GetServers(host)
	if err != nil {

		fmt.Println(err)
		log.Println("Database: Error al obtener los servidores del dominio.")
		return Domain{}, err
	}

	return domain, nil
}

func GetPreviousDomains() ([]Domain, error) {

	var previous []Domain

	database := ConnectDatabase()
	defer database.Close()

	queryDomain :=
		`SELECT *
		FROM domain;`

	rows, err := database.Query(queryDomain)
	defer rows.Close()

	if err != nil {

		fmt.Println(err)
		log.Println("Database: No se pudo obtener los dominios.")
		return []Domain{}, err
	}

	for rows.Next() {
		var domain Domain

		err := rows.Scan(
			&domain.Domain,
			&domain.Info.Servers_changed,
			&domain.Info.Ssl_grade,
			&domain.Info.Previous_ssl_grade,
			&domain.Info.Logo,
			&domain.Info.Title,
			&domain.Info.Is_down,
			&domain.Updated,
		)

		if err != nil {

			fmt.Println(err)
			log.Println("Database: Error al escanear los datos de los dominios.")
			return []Domain{}, err
		}

		domain.Info.Servers, err = GetServers(domain.Domain)
		if err != nil {

			fmt.Println(err)
			log.Println("Database: Error al obtener los servidores de los dominios.")
			return []Domain{}, err
		}

		previous = append(previous, domain)
	}

	return previous, nil
}

func UpdateDomain(domain string, changed bool, sslgrade string, previoussslgrade string, updated string) error {

	database := ConnectDatabase()
	defer database.Close()
	var flag bool

	queryDomain :=
		`UPDATE domain 
		SET (servers_changed, ssl_grade, previous_ssl_grade, updated) = ($1, $2, $3, $4)
		WHERE domain = $5 RETURNING true;`

	err := database.QueryRow(queryDomain,
		changed, sslgrade, previoussslgrade, updated, domain).Scan(&flag)

	if err != nil {

		fmt.Println(err)
		log.Println("Database: No se pudo actualizar el dominio")
		return err
	}

	return nil
}
