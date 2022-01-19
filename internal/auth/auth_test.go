package auth

import (
	"airbound/internal/config"
	"airbound/internal/ent/enums"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
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

func TestGenerateAccountVerificationLink(t *testing.T) {
	is := require.New(t)

	cfg := &config.Config{JwtSecret: "test-jwt-secret", APIBaseURL: "http://localhost:4000"}
	is.NotNil(cfg)

	userID := uuid.MustParse("cad9b058-f106-4e7e-8801-eb1321fdb8d7")

	link, err := GenerateAccountVerificationLink(userID, cfg)
	is.NotEmpty(link)

	t.Errorf("[[link]] - %s", link)
	if err != nil {
		t.Errorf("expected 'err' to be nil, instead got %s", err)
	}
}

func TestValidateAccountVerificationToken(t *testing.T) {
	is := require.New(t)

	cfg := &config.Config{JwtSecret: "test-jwt-secret", APIBaseURL: "http://localhost:4000", AccountVerificationTTL: time.Second * 1}
	is.NotNil(cfg)

	userID := uuid.MustParse("cad9b058-f106-4e7e-8801-eb1321fdb8d7")

	link, err := GenerateAccountVerificationLink(userID, cfg)
	is.NotEmpty(link)
	is.NoError(err)

	token := strings.Split(link, "?tk=")[1]
	token = strings.Replace(token, "@", ".", -1)

	// time.Sleep(time.Second * 2)

	id, err := ValidateAccountVerificationToken(token+"1234", cfg)
	is.NotEmpty(id)
	is.NoError(err)

	if id != userID {
		t.Errorf("expected <id> to be equal to <userID>, instead got '%s'", id)
	}
}
