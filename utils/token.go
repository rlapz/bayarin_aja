package utils

import (
	"errors"
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

func TokenValidate(token string, secret []byte) (int64, error) {
	tok, err := jwt.Parse(token, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("a given token cannot be parsed")
		}

		return secret, nil
	})

	if err != nil {
		return -1, err
	}

	if cl, ok := tok.Claims.(jwt.MapClaims); ok && tok.Valid {
		return int64(cl["id"].(float64)), nil
	}

	return -1, errors.New("invalid token")
}
