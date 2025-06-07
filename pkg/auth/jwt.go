package auth

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTConfig struct {
	Algorithm string // RS256 or HS256
	Secret    string // for HS256
	PublicKey string // path to PEM for RS256
}

type JWTValidator struct {
	algo    string
	secret  []byte
	pubKey  *rsa.PublicKey
}

func NewJWTValidator(cfg JWTConfig) (*JWTValidator, error) {
	v := &JWTValidator{algo: cfg.Algorithm}

	switch cfg.Algorithm {
	case "HS256":
		v.secret = []byte(cfg.Secret)
	case "RS256":
		keyData, err := ioutil.ReadFile(cfg.PublicKey)
		if err != nil {
			return nil, fmt.Errorf("failed to read public key: %w", err)
		}
		pub, err := jwt.ParseRSAPublicKeyFromPEM(keyData)
		if err != nil {
			return nil, fmt.Errorf("invalid RSA public key: %w", err)
		}
		v.pubKey = pub
	default:
		return nil, fmt.Errorf("unsupported JWT algorithm: %s", cfg.Algorithm)
	}

	return v, nil
}

func (v *JWTValidator) Validate(tokenStr string) (jwt.MapClaims, error) {
	tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")

	var keyFunc jwt.Keyfunc
	switch v.algo {
	case "HS256":
		keyFunc = func(t *jwt.Token) (interface{}, error) {
			return v.secret, nil
		}
	case "RS256":
		keyFunc = func(t *jwt.Token) (interface{}, error) {
			return v.pubKey, nil
		}
	}

	token, err := jwt.Parse(tokenStr, keyFunc, jwt.WithValidMethods([]string{v.algo}))
	if err != nil || !token.Valid {
		return nil, errors.New("invalid JWT")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid claims")
	}

	// optional: check expiry
	if exp, ok := claims["exp"].(float64); ok && time.Now().Unix() > int64(exp) {
		return nil, errors.New("token expired")
	}

	return claims, nil
}
