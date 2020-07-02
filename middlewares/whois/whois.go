package whois

import (
	"log"

	"github.com/likexian/whois-go"
	whoisparser "github.com/likexian/whois-parser-go"
)

func WhoisDomain(host string) ([]string, bool) {

	var whoisDomain []string
	var success bool

	response, err := whois.Whois(host)
	if err != nil {

		log.Println("Whois: Error con el dominio ingresado.")
		success = false
	} else {

		result, err := whoisparser.Parse(response)
		if err != nil {

			log.Println("Whois: No se pudieron obtener los datos del dominio.")
			success = false
		} else {
			success = true
			whoisDomain = append(whoisDomain, result.Administrative.Country)
			whoisDomain = append(whoisDomain, result.Administrative.Organization)
		}

	}

	return whoisDomain, success
}
