package twofa

import (
	"airbound/internal/config"
	"airbound/internal/log"
	"bytes"
	"encoding/base64"
	"image/png"

	"github.com/pquerna/otp/totp"
)

type TwoFaKeyPair struct {
	Secret string
	QRCode string
}

func GenerateKeyPair(email string, cfg *config.Config) (*TwoFaKeyPair, error) {
	logger := log.WithField("FunctionName", "GenerateKeyPair")

	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      cfg.TwoFaIssuer,
		AccountName: email,
		// Period:      uint(cfg.TwoFaTTL),
		// Digits:      otp.Digits(cfg.TwoFaLength),
	})
	if err != nil {
		logger.Error("error generating totp key - %s", err)
		return nil, err
	}

	var buf bytes.Buffer
	img, err := key.Image(200, 200)
	if err != nil {
		logger.Error("error coverting TOTP key to PNG - %s", err)
		return nil, err
	}

	if err := png.Encode(&buf, img); err != nil {
		logger.Error("error writing image to buffer - %s", err)
		return nil, err
	}

	kp := &TwoFaKeyPair{}
	b := buf.Bytes()

	qrcode := base64.StdEncoding.EncodeToString(b)

	kp.QRCode = qrcode
	kp.Secret = key.Secret()

	return kp, nil
}

func Validate2FaCode(code string, secret string) bool {
	return totp.Validate(code, secret)
}
