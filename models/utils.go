package models

import (
	"crypto/rand"
	"math/big"

	"golang.org/x/crypto/bcrypt"
)

// hash password using bcrypt
func HashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

func VerifyPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

// create a cryptographically secure 6-digit code
func GenerateResetCode() (string, error) {
	code := make([]byte, 6)
	for i := range 6 {
		num, err := rand.Int(rand.Reader, big.NewInt(10))
		if err != nil {
			return "", err
		}
		code[i] = byte(num.Int64()) + '0'
	}
	return string(code), nil
}
