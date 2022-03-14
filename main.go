package main

import (
	_ "github.com/joho/godotenv/autoload"
	"log"

	"blog-gin_golang_v177/config"
)

func main() {
	err := config.Routers.Run()
	if err != nil {
		log.Fatal(err)
	}
}
