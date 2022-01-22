package twofa

import (
	"airbound/internal/config"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestGenerateKeyPair(t *testing.T) {
	is := require.New(t)
	cfg := &config.Config{TwoFaIssuer: "Airbound", TwoFaTTL: time.Duration(time.Second * 45), TwoFaLength: 8}

	_, err := GenerateKeyPair("test@email.com", cfg)
	is.NoError(err)
}
