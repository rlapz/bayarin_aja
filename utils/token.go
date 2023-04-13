package utils

import (
	"time"

	"github.com/golang-jwt/jwt"
)

// id    : customer id
// secret: server secret
// exp   : expiration time
//
// return: generated token if success
func TokenGenerate(secret []byte, id int64, exp time.Duration) (string, error) {
	tok := jwt.NewWithClaims(jwt.SigningMethodHS256, newClaims(id, exp))
	signed, err := tok.SignedString(secret)
	if err != nil {
		return "", err
	}

	return signed, nil
}

func newClaims(id int64, exp time.Duration) jwt.MapClaims {
	return jwt.MapClaims{
		"id": id,
	}
}
