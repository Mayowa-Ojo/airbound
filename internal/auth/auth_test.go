package auth

import (
	"airbound/internal/config"
	"airbound/internal/ent/enums"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidatePassword(t *testing.T) {
	password := "jhfhgBj6hd#$%"

	if err := ValidatePassword(password); err != nil {
		t.Errorf("expected 'err' to be nil, instead got %s", err)
	}
}

func TestVerifyToken(t *testing.T) {
	is := require.New(t)

	cfg := &config.Config{JwtSecret: "test-jwt-secret"}
	token, err := GenerateToken(TokenData{Role: "ADMIN"}, cfg)
	is.NoError(err)

	tokenData, err := VerifyToken(token, cfg)
	is.NoError(err)

	expected := "ADMIN"
	if tokenData.Role != enums.Role(expected) {
		t.Errorf("expected role to be '%s', instead got '%s'", expected, tokenData.Role)
	}
}
