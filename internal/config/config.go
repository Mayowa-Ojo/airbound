package config

import (
	"log"
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	APIBaseURL             string        `envconfig:"API_BASE_URL"`
	Port                   int           `envconfig:"PORT"`
	DBDriver               string        `envconfig:"DB_DRIVER"`
	DBName                 string        `envconfig:"DB_NAME"`
	DBHost                 string        `envconfig:"DB_HOST"`
	DBUser                 string        `envconfig:"DB_USER"`
	DBPassword             string        `envconfig:"DB_PASSWORD"`
	DBPort                 int           `envconfig:"DB_PORT"`
	JwtSecret              string        `envconfig:"JWT_SECRET"`
	TokenTTL               time.Duration `envconfig:"TOKEN_TTL"`
	AWSRegion              string        `envconfig:"AWS_REGION"`
	AWSSESAccessKeyID      string        `envconfig:"AWS_SES_ACCESS_KEY_ID"`
	AWSSESSecretAccessKey  string        `envconfig:"AWS_SES_SECRET_ACCESS_KEY"`
	AWSSESARN              string        `envconfig:"AWS_SES_ARN"`
	EmailFromName          string        `envconfig:"EMAIL_FROM_NAME"`
	EmailFromAddress       string        `envconfig:"EMAIL_FROM_ADDRESS"`
	AccountVerificationTTL time.Duration `envconfig:"ACCOUNT_VERIFICATION_TTL"`
	TwoFaIssuer            string        `envconfig:"TWO_FA_ISSUER"`
	TwoFaTTL               time.Duration `envconfig:"TWO_FA_TTL"`
	TwoFaLength            int           `envconfig:"TWO_FA_LENGTH"`
	MaxAircraftAge         time.Duration `envconfig:"MAX_AIRCRAFT_AGE"`
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
