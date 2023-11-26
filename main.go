package main

import (
	"log"
	server "server/api"
)

func main() {
	err := server.RunServer(":4000")
	log.Fatal(err)
}
