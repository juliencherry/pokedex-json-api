package env

import (
	"github.com/joeshaw/envdecode"
	"github.com/joho/godotenv"
)

type Config struct {
	Port uint16 `env:"PORT,default=8081"`
}

func LoadEnvironment() (Config, error) {
	var cfg Config

	_ = godotenv.Load()

	err := envdecode.Decode(&cfg)
	if err != nil {
		return cfg, err
	}

	return cfg, nil
}
