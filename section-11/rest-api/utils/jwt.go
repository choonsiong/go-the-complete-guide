package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"time"
)

const secretKey = "my-secret-key"

func GenerateToken(email string, userID int) (string, error) {
	log.Println("GenerateToken -", email, userID)
	return jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userID": userID,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	}).SignedString([]byte(secretKey))
}

func VerifyToken(tokenString string) (string, int, error) {
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Make sure signing method is valid
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpected signing method")
		}

		return []byte(secretKey), nil
	})
	if err != nil {
		return "", 0, errors.New("could not parse token")
	}

	if !parsedToken.Valid {
		return "", 0, errors.New("invalid token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return "", 0, errors.New("invalid token claims")
	}

	email, ok := claims["email"].(string)
	log.Println("VerifyToken -", email)
	if !ok {
		return "", 0, errors.New("invalid email")
	}

	userID, ok := claims["userID"].(float64)
	log.Println("VerifyToken -", userID)
	if !ok {
		return "", 0, errors.New("invalid user id")
	}

	return email, int(userID), nil
}
