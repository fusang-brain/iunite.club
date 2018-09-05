package utils

import (
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"time"
	// "github.com/goadesign/goa/middleware/security/jwt"
)

// GenerateToken 生成Token
func GenerateToken(user string) (string, error) {
	// Create Token
	token := jwt.New(jwt.SigningMethodHS256)
	// // Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["iss"] = user + fmt.Sprintf("_%d", time.Now().Unix())
	claims["user"] = user
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	t, err := token.SignedString([]byte("c13be55b40cf9dacb8231156ff28d41e65c8b48b"))
	if err != nil {
		return "", err
	}
	return t, nil
}

// DecodeToken is a func get token from string
func DecodeToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}

		return []byte("c13be55b40cf9dacb8231156ff28d41e65c8b48b"), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, err
}
