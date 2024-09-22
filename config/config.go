package config

import (
	"log"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	AppPort     int `env:"APP_PORT" envDefault:"5000"`
	AppHttpPort int `env:"APP_HTTP_PORT" envDefault:"5001"`

	DBHost string `env:"DB_HOST" envDefault:"localhost"`
	DBPort int32  `env:"DB_PORT" envDefault:"5432"`
	DBName string `env:"DB_NAME,required"`
	DBUser string `env:"DB_USER,required"`
	DBPwd  string `env:"DB_PWD,required"`

	TgApiID   int32  `env:"TG_API_ID,required"`
	TgApiHash string `env:"TG_API_HASH,required"`
}

func LoadConfig() *Config {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		log.Fatalf("cannot parse initial ENV vars: %v", err)
	}
	return cfg
}
