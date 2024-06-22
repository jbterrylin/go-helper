package jwthelper

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
	BufferTime int64
	jwt.RegisteredClaims
	Data map[string]interface{}
}
