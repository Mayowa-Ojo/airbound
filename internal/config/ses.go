package config

import (
	logger "airbound/internal/log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

func NewSESSession(cfg *Config) *session.Session {
	sess, err := session.NewSession(&aws.Config{
		Region: &cfg.AWSRegion,
		Credentials: credentials.NewStaticCredentials(
			cfg.AWSSESAccessKeyID,
			cfg.AWSSESSecretAccessKey,
			"",
		),
	})
	if err != nil {
		logger.Fatal("[SES]: error creating new session - %s", err)
	}

	if _, err := sess.Config.Credentials.Get(); err != nil {
		logger.Fatal("[SES]: error fetching credentials value - %s", err)
	}

	logger.Info("[SES]: aws session created")

	return sess
}
