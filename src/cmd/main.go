package main

import (
	"crud-user/src/business/domain"
	"crud-user/src/business/usecase"
	"crud-user/src/handler/rest"
	sql "crud-user/src/lib/mysql"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("failed to load env from local file")
	}

	dbConfig := sql.Config{
		Host:     os.Getenv("DB_HOST"),
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		Database: os.Getenv("DB_NAME"),
		Port:     os.Getenv("DB_PORT"),
	}
	db := sql.Init(dbConfig)

	d := domain.Init(db)

	uc := usecase.Init(d)

	r := rest.Init(uc)

	r.Run()
}
