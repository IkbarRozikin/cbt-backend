package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var jwtSecret = []byte("your_jwt_secret_key")

func GenerateJWT(userID, roleID, schoolID uuid.UUID) (string, error) {
	claims := jwt.MapClaims{
		"user_id":   userID,
		"role_id":   roleID,
		"school_id": schoolID,
		"exp":       time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
