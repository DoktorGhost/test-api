package config

import (
	"github.com/caarlos0/env/v6"
	"log"
)

type Config struct {
	DBHost  string `env:"DB_HOST"`
	DBPort  string `env:"DB_PORT"`
	DBUser  string `env:"DB_USER"`
	DBPass  string `env:"DB_PASSWORD"`
	DBName  string `env:"DB_NAME"`
	ApiPort string `env:"API_PORT"`
	ApiHost string `env:"API_HOST"`
}

func ParseConfigServer() Config {
	// Чтение переменных окружения
	config := Config{}
	if err := env.Parse(&config); err != nil {
		log.Println(err)
	}
	return config
}
