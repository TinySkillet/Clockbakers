package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	m "github.com/TinySkillet/ClockBakers/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type apiHandler func(http.ResponseWriter, *http.Request)

func MiddlewareValidateUser(next apiHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token, err := getBearerToken(r)
		if err != nil {
			http.Error(w, "Not Authorized!", http.StatusUnauthorized)
			return
		}
		claims, err := verifyToken(token)
		if err != nil {
			http.Error(w, "Not allowed!"+err.Error(), http.StatusUnauthorized)
			return
		}

		query := r.URL.Query()
		uid := query.Get("uid")
		// ensures that users can only access their own data
		if uid != "" {
			if uid != claims.ID {
				http.Error(w, "Not allowed! UID mismatch", http.StatusUnauthorized)
				return
			}
		}

		if claims.Role == "customer" || claims.Role == "admin" {
			next(w, r)
			return
		}
		http.Error(w, "Not allowed! Invalid role!", http.StatusUnauthorized)
	}
}

func MiddlewareValidateAdmin(next apiHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token, err := getBearerToken(r)
		if err != nil {
			http.Error(w, "Not Authorized!", http.StatusUnauthorized)
			return
		}
		claims, err := verifyToken(token)
		if err != nil {
			http.Error(w, "Not allowed! "+err.Error(), http.StatusUnauthorized)
			return
		}
		if claims.Role != "admin" {
			http.Error(w, "Not allowed!", http.StatusUnauthorized)
			return
		}
		next(w, r)
	}
}

// getBearerToken retrieves the value from Authorization key in request header
func getBearerToken(r *http.Request) (string, error) {
	token := r.Header.Get("Authorization")
	if token == "" {
		return "", errors.New("No authorization token found!")
	}

	splits := strings.Split(token, " ")
	if len(splits) != 2 || splits[0] != "Bearer" {
		return "", errors.New("Invalid authorization token!")
	}
	return splits[1], nil
}

func getSigningKey() (string, error) {
	key, found := os.LookupEnv("SECRET_KEY")
	if !found {
		return "", errors.New("JWT signing key not found in the environment!")
	}
	return key, nil
}

// CreateToken creates a jwt token with the given id, email and role
func CreateToken(id uuid.UUID, email, role string) (string, error) {
	key, err := getSigningKey()
	if err != nil {
		return "", err
	}
	claims := m.JWTClaims{
		Email: email,
		Role:  role,
		RegisteredClaims: jwt.RegisteredClaims{
			ID: id.String(),
			// 12 hours validity
			ExpiresAt: &jwt.NumericDate{
				Time: time.Now().Add(time.Hour * 12).UTC(),
			},
			IssuedAt: &jwt.NumericDate{
				Time: time.Now().UTC(),
			},
			Issuer: "localhost:6000",
		},
	}
	jwt_token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	tokenString, err := jwt_token.SignedString([]byte(key))
	if err != nil {
		return "", errors.New("Error generating JSON web token!")
	}
	return tokenString, nil
}

// VerifyToken checks and verifies only the signing method, expiry date and claims structure of the token
func verifyToken(tokenString string) (*m.JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &m.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		// check signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		key, err := getSigningKey()
		return []byte(key), err
	})
	if err != nil {
		return nil, err
	}

	// extract claims correctly
	claims, ok := token.Claims.(*m.JWTClaims)
	if !ok || !token.Valid {
		return nil, errors.New("Invalid token claims!")
	}

	fmt.Printf("User: %s | Role: %s", claims.Email, claims.Role)
	return claims, nil
}
