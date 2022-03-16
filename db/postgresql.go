package db

import (
	"blog-gin_golang_v177/lib/env"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	// "gorm.io/gorm/logger"
)

func Postgresql() (*gorm.DB, error) {
	return gorm.Open(postgres.Open(dsn()), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Info),
	})
}

func dsn() string {
	host := "host=" + env.String("DATABASE_HOST", "localhost")
	port := "port=" + env.String("DATABASE_PORT", "5432")
	dbname := "dbname=" + env.String("DATABASE_NAME", "codedoct_gin_golang177")
	user := "user=" + env.String("DATABASE_USER", "postgres")
	password := "password=" + env.String("DATABASE_PASSWORD", "YpP1R7Z09q")
	return fmt.Sprintln(host, port, dbname, user, password)
}
