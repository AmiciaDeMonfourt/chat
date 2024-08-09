package jwt

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

var signedKey = []byte("key")

func GenerateToken(userid uint64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"user_id": userid})

	tokenString, err := token.SignedString(signedKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func CheckAuth(tokenStr string) error {
	return checkToken(tokenStr)
}

func ExtractUserId(tokenStr string) (uint64, error) {
	if err := checkToken(tokenStr); err != nil {
		return 0, err
	}

	token, err := jwt.Parse(tokenStr, keyFunc)
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return 0, errors.New("invalid token")
	}

	userId, exists := claims["user_id"].(float64)
	if !exists {
		return 0, err
	}

	return uint64(userId), nil
}

func checkToken(tokenStr string) error {
	token, err := jwt.Parse(tokenStr, keyFunc)
	if err != nil {
		return fmt.Errorf("failed to parse token: " + err.Error())
	}

	if !token.Valid {
		return fmt.Errorf("token is invalid")
	}

	return nil
}

func keyFunc(*jwt.Token) (interface{}, error) {
	return signedKey, nil
}
