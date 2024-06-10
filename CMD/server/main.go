package main

import (
	"api/internal/config"
	"api/internal/server"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	//conf := config.ParseConfigServer()

	conf := config.Config{
		Host:        "localhost",
		Port:        "8080",
		DatabaseDSN: "postgresql://admin:admin@localhost:5432/admin",
	}

	err := server.StartServer(&conf)

	if err != nil {
		panic(err)
	}
}
