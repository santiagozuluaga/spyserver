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
	Updated string     `json:"updated"`
}
