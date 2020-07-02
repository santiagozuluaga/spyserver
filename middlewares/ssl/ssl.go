package ssl

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/Santiagozh1998/SpyServer/middlewares/database"
	"github.com/valyala/fasthttp"
)

func SslDomain(host string) (database.SslInfo, bool) {

	var body []byte
	var result database.SslInfo
	var endpoints []database.Endpoints
	success := true

	url := fmt.Sprintf("https://api.ssllabs.com/api/v3/analyze?host=%s", host)
	_, body, err := fasthttp.Get(body, url)

	if err != nil {

		log.Println("SSL: Error al obtener los datos de Ssllabs.")
		success = false
	} else {

		json.Unmarshal(body, &result)

		if result.Status == "READY" && len(result.Endpoints) > 0 {

			for i := 0; i < len(result.Endpoints); i++ {

				if result.Endpoints[i].StatusMessage == "Ready" {

					endpoints = append(endpoints, result.Endpoints[i])
				}
			}

			if len(endpoints) == 0 {

				success = false
				log.Println("SSL: El servidor no tiene endpoints aun")
			} else {

				result.Endpoints = endpoints
			}

		} else {

			log.Println("SSL: No se pueden obtener los datos del dominio.")
			success = false
		}

	}

	return result, success
}
