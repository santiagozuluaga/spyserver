package middlewares

import (
	"github.com/Santiagozh1998/SpyServer/middlewares/database"
	"github.com/Santiagozh1998/SpyServer/middlewares/date"
	"github.com/Santiagozh1998/SpyServer/middlewares/scraping"
	"github.com/Santiagozh1998/SpyServer/middlewares/ssl"
	"github.com/Santiagozh1998/SpyServer/middlewares/whois"
)

var SSLGRADES = []string{"F", "E", "D", "C", "B", "A", "A+"}

func CheckServers(endpoints []database.Endpoints, whois []string) (int, []database.Server) {

	var servers []database.Server
	var server database.Server
	index := 6

	for i := 0; i < len(endpoints); i++ {

		server.Address = endpoints[i].IpAddress
		server.Ssl_grade = endpoints[i].Grade
		server.Country = whois[0]
		server.Owner = whois[1]

		for g := 0; g < len(SSLGRADES); g++ {

			if SSLGRADES[g] == endpoints[i].Grade {
				if index > g {
					index = g
				}
			}
		}

		servers = append(servers, server)

	}

	return index, servers
}

func CheckDomain(domain database.Domain) database.DomainInfo {

	flag, newDate := date.UpdateDate(domain.Updated)

	if flag == true {

		sslhost, flagssl := ssl.SslDomain(domain.Domain)
		whoishost, flagwhois := whois.WhoisDomain(domain.Domain)

		if flagssl && flagwhois {

			index, servers := CheckServers(sslhost.Endpoints, whoishost)

			domain.Info.Previous_ssl_grade = domain.Info.Ssl_grade
			domain.Info.Ssl_grade = SSLGRADES[index]

			return CheckUpdate(domain, servers, newDate)

		}

	}

	return domain.Info
}

func CheckUpdate(domain database.Domain, servers []database.Server, newDate string) database.DomainInfo {

	updateArray, insertArray, deleteArray := CheckArrays(servers, domain.Info.Servers)

	flagUpdate := len(updateArray) != 0
	flagInsert := len(insertArray) != 0
	flagDelete := len(deleteArray) != 0

	if flagUpdate || flagInsert || flagDelete {

		if flagDelete {

			for i := 0; i < len(deleteArray); i++ {

				database.DeleteServer(domain.Info.Servers[i].Address)
			}
		}

		if flagInsert {

			for i := 0; i < len(insertArray); i++ {

				database.InsertServer(domain.Domain, servers[i].Address, servers[i].Ssl_grade, servers[i].Country, servers[i].Owner)
			}
		}

		if flagUpdate {

			for i := 0; i < len(updateArray); i++ {

				database.UpdateServer(servers[i].Address, servers[i].Ssl_grade)
			}
		}

		database.UpdateDomain(domain.Domain, true, domain.Info.Ssl_grade, domain.Info.Previous_ssl_grade, newDate)
		domain.Updated = newDate
		domain.Info.Servers = servers
		domain.Info.Servers_changed = true
	}

	return domain.Info
}

func CheckArrays(servers []database.Server, domain []database.Server) ([]int, []int, []int) {

	var updateArray, insertArray, deleteArray []int

	for i := 0; i < len(servers); i++ {

		flagInsert := true
		flagUpdate := true

		for j := 0; j < len(domain); j++ {

			if servers[i].Address == domain[j].Address {

				flagInsert = false

				if servers[i].Ssl_grade == domain[j].Ssl_grade {

					flagUpdate = false
				}

				break
			}
		}

		if flagInsert {

			insertArray = append(insertArray, i)
		} else {

			if flagUpdate {

				updateArray = append(updateArray, i)
			}
		}
	}

	for i := 0; i < len(domain); i++ {

		flagDelete := true

		for j := 0; j < len(servers); j++ {

			if domain[i].Address == servers[j].Address {

				flagDelete = false
				break
			}
		}

		if flagDelete {

			deleteArray = append(deleteArray, i)
		}
	}

	return updateArray, insertArray, deleteArray
}

func CreateDomain(host string) database.DomainInfo {

	var domain database.Domain

	sslhost, flagssl := ssl.SslDomain(host)

	if flagssl {

		whoishost, flagwhois := whois.WhoisDomain(host)
		scrapinghost, flagscraping := scraping.ScrapingDomain(host)

		if flagwhois && flagscraping {

			domain.Domain = host
			dateNow := date.CreateDate()
			domain.Created = dateNow
			domain.Updated = dateNow

			domain.Info.Servers_changed = false
			domain.Info.Is_down = false
			domain.Info.Title = scrapinghost[0]
			domain.Info.Logo = scrapinghost[1]
			index, servers := CheckServers(sslhost.Endpoints, whoishost)
			domain.Info.Servers = servers
			domain.Info.Ssl_grade = SSLGRADES[index]
			domain.Info.Previous_ssl_grade = SSLGRADES[index]

			database.InsertDomain(domain)
		}

	}

	return domain.Info

}

func SearchDomain(host string) database.DomainInfo {

	domain, flag := database.GetDomain(host)

	if flag {

		return CheckDomain(domain)
	} else {

		return CreateDomain(host)
	}

}

func PreviousDomains() []database.Domain {

	var domains []database.Domain

	domains = database.GetPreviousDomains()

	return domains
}
