package config

import (
	"github.com/caarlos0/env/v6"
	"log"
)

type Config struct {
	Host        string `env:"HOST"`
	Port        string `env:"PORT"`
	DatabaseDSN string `env:"DATABASE_DSN"`
}

func ParseConfigServer() *Config {
	// Чтение переменных окружения
	config := &Config{}
	if err := env.Parse(&config); err != nil {
		log.Println(err)
	}

	return config
}
