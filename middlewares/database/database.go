package database

import (
	"log"

	_ "github.com/lib/pq"
)

func InsertDomain(domain Domain) {

	database := ConnectDatabase()
	defer database.Close()
	var response string

	queryDomain :=
		`INSERT INTO domain 
			(domain, servers_changed, ssl_grade, previous_ssl_grade, logo, title, is_down, created, updated)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING domain.domain;`

	queryServers :=
		`INSERT INTO
		server (domain, address, ssl_grade, country, owner)
		VALUES ($1, $2, $3, $4, $5) RETURNING domain;`

	err := database.QueryRow(queryDomain,
		domain.Domain,
		domain.Info.Servers_changed,
		domain.Info.Ssl_grade,
		domain.Info.Previous_ssl_grade,
		domain.Info.Logo,
		domain.Info.Title,
		domain.Info.Is_down,
		domain.Created,
		domain.Updated).Scan(&response)

	if err != nil {

		log.Println("Database: No se pudo ingresar el dominio")
	} else {

		for i := 0; i < len(domain.Info.Servers); i++ {

			err := database.QueryRow(queryServers,
				response,
				domain.Info.Servers[i].Address,
				domain.Info.Servers[i].Ssl_grade,
				domain.Info.Servers[i].Country,
				domain.Info.Servers[i].Owner).Scan(&response)

			if err != nil {

				log.Println("Database: No se pudo ingresar los servidores")
			}
		}
	}

}

func GetDomain(host string) (Domain, bool) {

	database := ConnectDatabase()
	defer database.Close()

	var domain Domain
	var server Server
	flag := false

	queryDomain :=
		`SELECT *
		FROM domain 
		WHERE domain = $1;`

	queryServers :=
		`SELECT address, ssl_grade, country, owner
		FROM server
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
		&domain.Created,
		&domain.Updated,
	)

	if err != nil {

		log.Println("Database: No se pudo obtener el dominio")
	} else {

		rows, err := database.Query(queryServers, host)
		defer rows.Close()

		if err != nil {
			log.Println("Database: No se pudo obtener los datos de los servidores.")
		} else {

			for rows.Next() {

				err := rows.Scan(
					&server.Address,
					&server.Ssl_grade,
					&server.Country,
					&server.Owner,
				)

				if err == nil {

					flag = true
					domain.Info.Servers = append(domain.Info.Servers, server)
				} else {
					log.Println("Database: Error al escanear los datos de los servidores.")
					flag = false
				}

			}
		}

	}

	return domain, flag
}

func GetPreviousDomains() []Domain {

	var previous []Domain

	database := ConnectDatabase()
	defer database.Close()

	queryDomain :=
		`SELECT *
		FROM domain;`

	queryServers :=
		`SELECT address, ssl_grade, country, owner
		FROM server
		WHERE domain = $1;`

	rows, err := database.Query(queryDomain)
	defer rows.Close()

	if err != nil {

		log.Println("Database: No se pudo obtener los dominios.")
	} else {

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
				&domain.Created,
				&domain.Updated,
			)

			if err != nil {

				log.Println("Database: Error al escanear los datos de los dominios.")
			} else {

				rows, err := database.Query(queryServers, domain.Domain)
				defer rows.Close()

				if err != nil {

					log.Println("Database: No se pudieron obtener los servidores.")
				} else {

					for rows.Next() {

						var server Server

						err := rows.Scan(
							&server.Address,
							&server.Ssl_grade,
							&server.Country,
							&server.Owner,
						)

						if err != nil {

							log.Println("Database: Error al escanear los datos de los servidores.")
						} else {

							domain.Info.Servers = append(domain.Info.Servers, server)
						}
					}

					previous = append(previous, domain)
				}
			}

		}
	}

	return previous
}

func UpdateDomain(domain string, changed bool, sslgrade string, previoussslgrade string, updated string) {

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

		log.Println("Database: No se pudo actualizar el dominio")
	}
}

func UpdateServer(address string, ssl_grade string) {

	database := ConnectDatabase()
	defer database.Close()
	var flag bool

	queryServer :=
		`UPDATE server 
		SET ssl_grade = $1
		WHERE address = $2;`

	err := database.QueryRow(queryServer,
		ssl_grade, address).Scan(&flag)

	if err != nil {

		log.Println("Database: No se pudo actualizar el servidor")
	}

}

func DeleteServer(address string) {

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

		log.Println("Database: No se pudo eliminar el servidor")
	}
}

func InsertServer(domain string, address string, ssl_grade string, country string, owner string) {

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

		log.Println("Database: No se pudo ingresar el servidor")
	}
}
