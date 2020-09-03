package routes

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/Santiagozh1998/SpyServer/src/database"
	"github.com/valyala/fasthttp"

	"github.com/buaazp/fasthttprouter"
)

func SearchHandler(ctx *fasthttp.RequestCtx) {

	ctx.SetContentType("application/json")

	host := fmt.Sprintf("%s", ctx.UserValue("host"))

	json.NewEncoder(ctx).Encode(CheckDomain(host))
}

func PreviousHandler(ctx *fasthttp.RequestCtx) {

	ctx.SetContentType("application/json")

	domains, err := database.GetPreviousDomains()
	if err != nil {
		fmt.Println(err)
		log.Println("Database: No se pudo ingresar el dominio")
	}
	json.NewEncoder(ctx).Encode(domains)
}

func AppRouter() *fasthttprouter.Router {

	router := fasthttprouter.New()
	router.GET("/api/servers/search/:host", SearchHandler)
	router.GET("/api/servers/previous", PreviousHandler)

	return router
}
