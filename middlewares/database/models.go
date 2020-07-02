package database

//DATABASE
type Server struct {
	Address   string `json:"address"`
	Ssl_grade string `json:"ssl_grade"`
	Country   string `json:"country"`
	Owner     string `json:"owner"`
}

type DomainInfo struct {
	Servers            []Server `json:"servers"`
	Servers_changed    bool     `json:"servers_changed"`
	Ssl_grade          string   `json:"ssl_grade"`
	Previous_ssl_grade string   `json:"previous_ssl_grade"`
	Logo               string   `json:"logo"`
	Title              string   `json:"title"`
	Is_down            bool     `json:"is_down"`
}

type Domain struct {
	Domain  string     `json:"domain"`
	Info    DomainInfo `json:"info"`
	Created string     `json:"created"`
	Updated string     `json:"updated"`
}

//APIS
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
