package main

import (
	"log"
	"os"

	"github.com/Santiagozh1998/SpyServer/src/routes"
	"github.com/valyala/fasthttp"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	router := routes.AppRouter().Handler

	log.Println("Server running in port: " + port)
	fasthttp.ListenAndServe(":"+port, router)
}
