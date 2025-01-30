package main

import (
	"log"

	"mpozdal/remitly/cmd/api"
)

func main() {

	server := api.NewAPIServer(":8080")
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
