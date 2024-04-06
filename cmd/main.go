package main

import (
	"log"

	"github.com/riju-stone/go-rest-api/cmd/api"
)

func main() {
	server := api.CreateServer(":8080", nil)
	if err := server.RunServer(); err != nil {
		log.Fatal(err)
	}
}
