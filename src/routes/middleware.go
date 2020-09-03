package routes

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Santiagozh1998/SpyServer/src/database"
	"github.com/Santiagozh1998/SpyServer/src/utils"
)

func SmallerSSLGRADE(grades []string) string {

	smaller := grades[0]

	if len(grades) > 1 {

		for i := 1; i < len(grades); i++ {

			if smaller < grades[i] && grades[i] != "A+" {

				smaller = grades[i]
			}
		}
	}

	return smaller
}

func CreateDomain(host string) (database.DomainInfo, error) {

	scraping, err := utils.ScrapingDomain(host)
	if err != nil {

		fmt.Println(err)
		log.Println("Database: No se pudo ingresar el dominio")
		return database.DomainInfo{}, err
	}

	servers, sslgrade, err := CreateServers(host)
	if err != nil {

		fmt.Println(err)
		log.Println("Database: No se pudo ingresar el dominio")
		return database.DomainInfo{}, err
	}

	database.InsertDomain(host, false, sslgrade, sslgrade, scraping[1], scraping[0], false, utils.GetTime())

	for i := 0; i < len(servers); i++ {

		database.InsertServer(host, servers[i].Address, servers[i].Ssl_grade, servers[i].Country, servers[i].Owner)
	}

	return database.DomainInfo{Servers: servers, Servers_changed: false, Previous_ssl_grade: sslgrade, Ssl_grade: sslgrade, Logo: scraping[1], Title: scraping[0], Is_down: false}, nil
}

func CreateServers(host string) ([]database.Server, string, error) {

	servers := []database.Server{}
	sslgrades := []string{}

	ssl, err := utils.SslTestFasthttp(host)
	if err != nil {

		fmt.Println(err)
		log.Println("Database: No se pudo ingresar el dominio")
		return []database.Server{}, "", err
	}

	for i := 0; i < len(ssl.Endpoints); i++ {

		whois, err := utils.WhoisDomain(ssl.Endpoints[i].IpAddress)
		if err != nil {

			fmt.Println(err)
			log.Println("Database: No se pudo ingresar el dominio")
			return []database.Server{}, "", err
		}

		sslgrades = append(sslgrades, ssl.Endpoints[i].Grade)

		servers = append(servers, database.Server{Address: ssl.Endpoints[i].IpAddress, Ssl_grade: ssl.Endpoints[i].Grade, Country: whois.Country, Owner: whois.Owner})
	}

	smaller := SmallerSSLGRADE(sslgrades)

	return servers, smaller, nil
}

func CheckDomain(host string) database.DomainInfo {

	domain, err := database.GetDomain(host)
	if err == sql.ErrNoRows {

		newdomain, err := CreateDomain(host)
		if err != nil {

			fmt.Println(err)
			log.Println("Database: No se pudo ingresar el dominio")
		}

		return newdomain
	}

	flag, newdate := utils.UpdateDate(domain.Updated)
	if flag == true {

		servers, sslgrade, err := CreateServers(host)
		if err != nil {

			fmt.Println(err)
			log.Println("Database: No se pudo ingresar el dominio")
		}

		flag := checkServers(host, domain.Info.Servers, servers)
		if flag == true {

			database.UpdateDomain(host, true, sslgrade, domain.Info.Ssl_grade, newdate)

			domain.Info.Previous_ssl_grade = domain.Info.Ssl_grade
			domain.Info.Ssl_grade = sslgrade
			domain.Info.Servers_changed = true
			domain.Info.Servers = servers
		}
	}

	return domain.Info

}

func checkServers(host string, current []database.Server, news []database.Server) bool {

	updated := false

	for i := 0; i < len(current); i++ {

		flag := false

		for j := 0; j < len(news); j++ {

			if current[i].Address == news[j].Address {

				flag = true

				if current[i].Ssl_grade != news[j].Ssl_grade {

					updated = true
					database.UpdateServer(news[j].Address, news[j].Ssl_grade)
				}
			}
		}

		if flag == false {

			updated = true
			database.DeleteServer(current[i].Address)
		}
	}

	for i := 0; i < len(news); i++ {

		flag := true

		for j := 0; j < len(current); j++ {

			if current[j].Address == news[i].Address {
				flag = false
			}
		}

		if flag == true {
			updated = true
			database.InsertServer(host, news[i].Address, news[i].Ssl_grade, news[i].Country, news[i].Owner)
		}
	}

	log.Println(current, news)

	return updated
}
