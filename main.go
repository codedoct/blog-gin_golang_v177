package main

import (
	"log"

	"blog-gin_golang_v177/config"
)

func main() {
	err := config.Routers.Run()
	if err != nil {
		log.Fatal(err)
	}
}
