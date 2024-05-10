package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secret = []byte("secret")

func GenerateToken(userId int64, email string) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"email":  email,
		"exp":    time.Now().Add(time.Hour * 24 * 365).Unix(),
	}).SignedString(secret)
}

func VerifyToken(token string) error {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpected signing method")
		}

		return secret, nil
	})

	if err != nil {
		return err
	}

	if !parsedToken.Valid {
		return errors.New("invalid token")
	}

	_, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return errors.New("invalid claims")
	}

	// email := claims["email"].(string)
	// userId := claims["userId"].(int64)

	return nil
}
