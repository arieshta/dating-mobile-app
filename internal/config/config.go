package config

import (
	"fmt"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

type Config struct {
	Port string `env:"PORT" envDefault:"8080"`

	JWT_SECRET string `env:"JWT_SECRET" envDefault:"secret"`
	JWT_expire string `env:"JWT_EXPIRED" envDefault:"5m"`

	MongoURI string `env:"MONGO_URI" envDefault:"mongodb://localhost:27017"`
	MongoDB  string `env:"MONGO_DB" envDefault:"dating_app"`
}

func LoadEnv() *Config {
	cfg := Config{}
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err, ", getting config from ENV")
	}
	err = env.Parse(&cfg)
	if err != nil {
		fmt.Println(err, ", getting config from ENV")
	}
	return &cfg
}
