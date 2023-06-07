package main

import (
	"log"

	"base/src/server"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}

	if err := server.StartServer(); err != nil {
		log.Fatal(err)
	}
}