package main

import (
	"log"
	"os"

	"github.com/cameronMcConnell/RoundRobin-Go/lib"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	serverAddress := os.Getenv("SERVER_ADDRESS")

	roundRobinServer, err := lib.NewRoundRobinServer(serverAddress)
	if err != nil {
		log.Fatal(err)
	}
	
	err = roundRobinServer.StartServer()
	if err != nil {
		log.Fatal(err)
	}
}