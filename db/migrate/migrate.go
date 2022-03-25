package main

import (
	"fmt"
	"log"

	"blog-gin_golang_v177/lib/env"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	m, err := migrate.New(
		"file://db/migrate/files",
		fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", env.String("DATABASE_USER", "postgres"), env.String("DATABASE_PASSWORD", "YpP1R7Z09q"), env.String("DATABASE_HOST", "localhost"), env.String("DATABASE_PORT", "5432"), env.String("DATABASE_NAME", "base_code")),
	)
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil && err.Error() != "no change" {
		log.Fatal(err)
	}
}
