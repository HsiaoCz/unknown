package api

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const token_expire = time.Hour * 24 * 30

var token_secert = []byte("xiaofanyi")

// get token
func GenJWT() (string, error) {
	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(token_expire)),
		Issuer:    "HsiaoCz",
	}

	// get jwt
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(token_secert)
}

// parse token
func PasreJWT(tokenString string) bool {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return token_secert, nil
	})

	if err != nil {
		return false
	}
	return token.Valid
}
