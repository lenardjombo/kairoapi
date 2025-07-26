package jwt

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"github.com/lenardjombo/kairoapi/models"
)

var jwtSecret []byte

func InitJWT() error {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("Warning: Failed to load .env file (might be OK if using system env vars): %v\n", err)
	}

	secretString := os.Getenv("JWT_SECRET")
	if secretString == "" {
		return fmt.Errorf("JWT_SECRET environment variable is not set")
	}

	jwtSecret = []byte(secretString)
	return nil
}

// GenerateToken generates a new JWT token for a given user.
func GenerateToken(userID, username string) (string, error) {
	if len(jwtSecret) == 0 {
		return "", fmt.Errorf("JWT secret not initialized. Call InitJWT() first")
	}

	claims := &models.UserClaims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)), // Token expires in 24 hours
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "kairoapi", // Application's name
			Subject:   userID,
			Audience:  []string{"users"},
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}
	return tokenString, nil
}

func ValidateToken(tokenString string) (*models.UserClaims, error) {
	if len(jwtSecret) == 0 {
		return nil, fmt.Errorf("JWT secret not initialized. Call InitJWT() first")
	}

	claims := &models.UserClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil 
	})

	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	if !token.Valid {
		return nil, fmt.Errorf("token is not valid")
	}

	return claims, nil
}