package config

import (
	"log"
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Port       int           `envconfig:"PORT"`
	DBDriver   string        `envconfig:"DB_DRIVER"`
	DBName     string        `envconfig:"DB_NAME"`
	DBHost     string        `envconfig:"DB_HOST"`
	DBUser     string        `envconfig:"DB_USER"`
	DBPassword string        `envconfig:"DB_PASSWORD"`
	DBPort     int           `envconfig:"DB_PORT"`
	JwtSecret  string        `envconfig:"JWT_SECRET"`
	TokenTTL   time.Duration `envconfig:"TOKEN_TTL"`
}

func LoadConfig() *Config {
	var config Config

	if err := godotenv.Load(); err != nil {
		log.Fatalf("[config]: error loading env file %s", err)
	}

	if err := envconfig.Process("", &config); err != nil {
		log.Fatalf("[config]: error populating env vars %s", err)
	}

	return &config
}
