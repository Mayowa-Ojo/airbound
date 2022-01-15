package auth

import (
	"airbound/internal/config"
	"airbound/internal/ent/enums"
	"airbound/internal/errors"
	"airbound/repository/actors"
	"crypto/rand"
	"crypto/subtle"
	"encoding/json"
	"fmt"
	"regexp"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/scrypt"
)

type TokenData struct {
	User  *actors.User
	Role  enums.Role
	Token string
}

const (
	saltLength = 32
	hashLength = 64
)

var (
	passwordConstraints = map[string]*regexp.Regexp{
		"length":      regexp.MustCompile(".{8,}"),
		"lowercase":   regexp.MustCompile(".*?[a-z]"),
		"uppercase":   regexp.MustCompile(".*?[A-Z]"),
		"number":      regexp.MustCompile(".*?[0-9]"),
		"specialChar": regexp.MustCompile(".*?[#?!@$ %^&*-]"),
	}
)

func ValidatePassword(s string) error {
	for _, v := range passwordConstraints {
		if !v.MatchString(s) {
			return errors.ErrInvalidPassword
		}
	}

	return nil
}

func VerifyPassword(password string, hash, salt []byte) error {
	h, err := generateHash(salt, password)
	if err != nil {
		return err
	}

	if subtle.ConstantTimeCompare(h, hash) != 1 {
		return errors.ErrIncorrectPassword
	}

	return nil
}

func HashPassword(password string) (hash, salt []byte, err error) {
	salt = generateSalt()
	hash, err = generateHash(salt, password)

	return
}

func generateSalt() []byte {
	salt := make([]byte, saltLength)

	_, _ = rand.Read(salt)

	return salt
}

func generateHash(s []byte, v string) ([]byte, error) {
	hash, err := scrypt.Key([]byte(v), s, 16384, 8, 1, hashLength)
	if err != nil {
		return nil, err
	}

	return hash, nil
}

func GenerateToken(data TokenData, cfg *config.Config) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["data"] = data
	claims["exp"] = time.Now().Add(cfg.TokenTTL).Unix()

	t, err := token.SignedString([]byte(cfg.JwtSecret))
	if err != nil {
		return "", err
	}

	return t, nil
}

func VerifyToken(t string, cfg *config.Config) (*TokenData, error) {
	token, err := jwt.ParseWithClaims(t, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %s", token.Header["alg"])
		}
		return []byte(cfg.JwtSecret), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid auth token")
	}

	var data *TokenData
	b, err := json.Marshal(claims["data"])
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(b, &data); err != nil {
		return nil, err
	}

	return data, nil
}
