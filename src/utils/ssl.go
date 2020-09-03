package utils

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/valyala/fasthttp"
)

type SslInfo struct {
	Host            string      `json:"host"`
	Port            int         `json:"port"`
	Protocol        string      `json:"protocol"`
	IsPublic        bool        `json:"isPublic"`
	Status          string      `json:"status"`
	StartTime       int         `json:"startTime"`
	TestTime        int         `json:"testTime"`
	EngineVersion   string      `json:"engineVersion"`
	CriteriaVersion string      `json:"criteriaVersion"`
	Endpoints       []Endpoints `json:"endpoints"`
}

type Endpoints struct {
	IpAddress         string `json:"ipAddress"`
	ServerName        string `json:"serverName"`
	StatusMessage     string `json:"statusMessage"`
	Grade             string `json:"grade"`
	GradeTrustIgnored string `json:"gradeTrustIgnored"`
	HasWarnings       bool   `json:"hasWarnings"`
	IsExceptional     bool   `json:"isExceptional"`
	Progress          int    `json:"progress"`
	Duration          int    `json:"duration"`
	Delegation        int    `json:"delegation"`
}

func SslTestFasthttp(host string) (SslInfo, error) {

	var result SslInfo

	url := fmt.Sprintf("https://api.ssllabs.com/api/v3/analyze?host=%s", host)
	_, body, err := fasthttp.Get(nil, url)
	if err != nil {

		fmt.Println(err)
		return result, err
	}

	json.Unmarshal(body, &result)

	if result.Status == "READY" && len(result.Endpoints) > 0 {

		for i := 0; i < len(result.Endpoints); i++ {

			if result.Endpoints[i].StatusMessage != "Ready" {

				return result, errors.New("NOT READY")
			}
		}

	} else {

		return result, errors.New("NOT READY")
	}

	return result, nil
}
