package main

import (
	"api/internal/config"
	"api/internal/server"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	conf := config.ParseConfigServer()

	err = server.StartServer(&conf)

	if err != nil {
		panic(err)
	}
}
