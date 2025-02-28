package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const secretKey = "my-secret-key"

func GenerateToken(email string, userID int) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userID": userID,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	}).SignedString([]byte(secretKey))
}
