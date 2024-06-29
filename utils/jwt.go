package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// For best practices take this from env file or valut or secrets.
const tokenGenValSecretKey = "someSecretKey"

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"expiry": time.Now().Add(time.Hour * 24).Unix(),
	})
	return token.SignedString([]byte(tokenGenValSecretKey))
}

func VerifyToken(token string) (*jwt.Token, error) {
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("signing method did not match")
		}

		return []byte(tokenGenValSecretKey), nil
	})

	if err != nil {
		return nil, errors.New("could not parse token")
	}

	validToken := parsedToken.Valid
	if !validToken {
		return nil, errors.New("invalid Token")
	}

	return parsedToken, nil
}

func GetDetailsFromJWT(token *jwt.Token) (map[string]interface{}, error) {

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return nil, errors.New("invalid token claims")
	}

	email, _ := claims["email"].(string)
	userId, _ := claims["userId"].(float64)

	data := make(map[string]interface{})
	data["email"] = email
	data["userId"] = int64(userId)
	return data, nil
}
