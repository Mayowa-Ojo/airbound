package config

import (
	"airbound/internal/log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

func NewSQSSession(cfg *Config) *session.Session {
	logger := log.WithField(string(log.LogFieldFunctionName), "<Config>NewSQSSession")

	sess, err := session.NewSession(&aws.Config{
		Region: &cfg.AWSRegion,
		Credentials: credentials.NewStaticCredentials(
			cfg.AWSSQSAccessKeyID,
			cfg.AWSSQSSecretAccessKey,
			"",
		),
	})
	if err != nil {
		logger.Fatal("error creating new session - %s", err)
	}

	if _, err := sess.Config.Credentials.Get(); err != nil {
		logger.Fatal("error fetching credentials value - %s", err)
	}

	logger.Info("[SQS]: aws session created")
	return sess
}
