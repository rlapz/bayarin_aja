package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

type Token struct {
	TokenString string
	ExpiresIn   time.Duration
}

// id    : customer id
// secret: server secret
// exp   : expiration time
//
// return: generated token and expires_in if success
func TokenGenerate(secret []byte, id int64, exp time.Duration) (Token, error) {
	claims := newClaims(id, exp)
	tok := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := tok.SignedString(secret)
	if err != nil {
		return Token{}, err
	}

	return Token{
		TokenString: signed,
		ExpiresIn:   time.Duration(claims["exp"].(int64)),
	}, nil
}

func newClaims(id int64, exp time.Duration) jwt.MapClaims {
	return jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(exp * time.Second).Unix(),
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
