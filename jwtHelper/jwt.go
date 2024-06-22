package jwthelper

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWT struct {
	SigningKey      []byte
	SigningMethod   jwt.SigningMethodHMAC
	BufferDuration  time.Duration
	ExpiredDuration time.Duration
	Issuer          string
	Audience        string
}

func NewJWT(signingKey []byte, SigningMethod jwt.SigningMethodHMAC) *JWT {
	return &JWT{
		SigningKey:    []byte(signingKey),
		SigningMethod: SigningMethod,
	}
}
