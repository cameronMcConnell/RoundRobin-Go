package main

import (
	"log"
	"github.com/cameronMcConnell/RoundRobin-Go/lib"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	
}