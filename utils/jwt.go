package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secret = "superman"

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":      email,
		"userId":     userId,
		"expireTime": time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(secret))
}

func VerifyToken(tokenString string) (int64, error) {
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("token signature does not match")
		}

		return []byte(secret), nil
	})

	if err != nil {
		return 0, err
	}

	if !parsedToken.Valid {
		return 0, fmt.Errorf("invalid token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid claim type")
	}

	userId := int64(claims["userId"].(float64))
	return userId, nil
}
