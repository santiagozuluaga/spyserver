package utils

import (
	"fmt"
	"os/exec"
	"strings"
)

type WhoisData struct {
	Owner   string
	Country string
}

func GetOwnerCountry(whois string) WhoisData {

	var data WhoisData

	linesWhois := strings.Split(whois, "\n")

	for i := 0; i < len(linesWhois); i++ {

		lineToLower := strings.ToLower(linesWhois[i])

		if strings.Contains(lineToLower, "orgname") || strings.Contains(lineToLower, "owner") || strings.Contains(lineToLower, "org-name") {

			data.Owner = strings.TrimSpace(strings.Split(linesWhois[i], ":")[1])
		}

		if strings.Contains(lineToLower, "country") {

			data.Country = strings.TrimSpace(strings.Split(linesWhois[i], ":")[1])
		}

		if data.Country != "" && data.Owner != "" {
			break
		}
	}

	return data
}

func WhoisDomain(host string) (WhoisData, error) {

	result, err := exec.Command("whois", host).Output()
	if err != nil {
		fmt.Println(err)
		return WhoisData{}, err
	}

	return GetOwnerCountry(string(result)), nil
}
