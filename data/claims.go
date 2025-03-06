package models

import (
	"github.com/golang-jwt/jwt/v5"
)

type JWTClaims struct {
	Email string
	Role  string
	jwt.RegisteredClaims
}
