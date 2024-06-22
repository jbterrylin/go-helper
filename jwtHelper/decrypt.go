package jwthelper

import (
	"github.com/golang-jwt/jwt/v5"
)

func (j *JWT) ParseToken(tokenString string) (claims Claims, err error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return
	}
	if token == nil {
		err = ErrTokenInvalid
		return
	}
	if !token.Valid {
		err = ErrTokenInvalid
		return
	}

	var ok bool
	claims, ok = token.Claims.(Claims)
	if !ok {
		err = ErrTokenInvalid
		return
	}

	return
}
