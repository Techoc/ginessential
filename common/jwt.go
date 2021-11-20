package common

import (
	"github.com/golang-jwt/jwt"
	"github.com/techoc/ginessential/model"
	"time"
)

var jwtKey = []byte("a_secret_key")

type Claims struct {
	UserId uint `json:"userId"`
	jwt.StandardClaims
}

func ReleaseToken(user model.User) (string, error) {
	//token有限期
	expireTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "techoc",
			Subject:   "user token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ParseToken(tokenString string) (token *jwt.Token, claims *Claims, err error) {
	claims = &Claims{}
	token, err = jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return
	}
	return token, claims, err
}
