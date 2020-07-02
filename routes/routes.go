package routes

import (
	"encoding/json"
	"fmt"

	"github.com/Santiagozh1998/SpyServer/middlewares"
	"github.com/valyala/fasthttp"

	"github.com/buaazp/fasthttprouter"
)

func CORS(next fasthttp.RequestHandler) fasthttp.RequestHandler {

	return func(ctx *fasthttp.RequestCtx) {

		ctx.Response.Header.Set("Access-Control-Allow-Credentials", "true")
		ctx.Response.Header.Set("Access-Control-Allow-Headers", "Authorization")
		ctx.Response.Header.Set("Access-Control-Allow-Methods", "HEAD,GET,POST,PUT,DELETE,OPTIONS")
		ctx.Response.Header.Set("Access-Control-Allow-Origin", "*")
		next(ctx)
	}
}

func SearchHandler(ctx *fasthttp.RequestCtx) {

	host := fmt.Sprintf("%s", ctx.UserValue("host"))

	domain := middlewares.SearchDomain(host)

	ctx.SetContentType("application/json")
	json.NewEncoder(ctx).Encode(domain)
}

func PreviousHandler(ctx *fasthttp.RequestCtx) {

	domains := middlewares.PreviousDomains()

	ctx.SetContentType("application/json")
	json.NewEncoder(ctx).Encode(domains)
}

func AppRouter() *fasthttprouter.Router {

	router := fasthttprouter.New()
	router.GET("/api/servers/search/:host", SearchHandler)
	router.GET("/api/servers/previous", PreviousHandler)

	return router
}
