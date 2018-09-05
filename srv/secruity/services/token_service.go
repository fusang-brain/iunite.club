package services

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/iron-kit/go-ironic/micro-assistant"
	pb "iunite.club/srv/user-srv/proto/user"
	"time"
)

type CustomClaims struct {
	UserID string
	jwt.StandardClaims
}

type TokenServicer interface {
	Decode(tokenString string) (*CustomClaims, error)
	Encode(user *pb.User) (string, error)
}

var privateSalt = []byte("c13be55b40cf9dacb8231156ff28d41e65c8b48b")

type TokenService struct {
	assistant.Service
}

/*
Decode is parse string to CustomClaims struct
*/
func (srv *TokenService) Decode(tokenString string) (*CustomClaims, error) {
	t, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return privateSalt, nil
	})

	if claims, ok := t.Claims.(*CustomClaims); ok && t.Valid {
		return claims, nil
	}

	return nil, err
}

/*
Encode is a function to generate a token
*/
func (srv *TokenService) Encode(user *pb.User, expireDay int64) (string, error) {
	if expireDay == 0 {
		expireDay = 7
	}
	now := time.Now()
	expireTime := now.Add(time.Hour * 24 * 7).Unix()

	claims := CustomClaims{
		UserID: user.ID,
		StandardClaims: jwt.StandardClaims{
			Issuer:    "kit.iron.srv.secruity", // 令牌签发者
			ExpiresAt: expireTime,              // 过期时间
			IssuedAt:  now.Unix(),              // 签发时间
		},
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	return jwtToken.SignedString(privateSalt)
}
