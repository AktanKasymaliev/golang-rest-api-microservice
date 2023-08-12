package main

import (
	"log"
	_ "base/docs"
	"base/internal/server"

	"github.com/joho/godotenv"
)

// @title           Golang REST API
// @version         1.0
// @description     A rest API service in Go using Gin framework.
// @termsOfService  https://tos.santoshk.dev

// @contact.name   	Aktan Kasymaliev
// @contact.url    	https://t.me/aktankasymaliev
// @contact.email  	akttan04@gmail.com

// @license.name  	GIN 2.0
// @license.url   	http://www.apache.org/licenses/LICENSE-2.0.html

// @host      		localhost:8000
// @BasePath  		/api/v1
func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}
	
	if err := server.StartServer(); err != nil {
		log.Fatal(err)
	}
}