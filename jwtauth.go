package main

import "os"
import "github.com/dgrijalva/jwt-go"
import "time"

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

type Claims struct {
	UserID  int    `json:"user_id"`
	UserTel string `json:"user_tel"`
	jwt.StandardClaims
}

func GenerateToken(id int, tel string) (string, error) {
	now := time.Now()
	exp := now.Add(6 * time.Hour)

	claims := Claims{
		id,
		tel,
		jwt.StandardClaims{
			ExpiresAt: exp.Unix(),
			Issuer:    "",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(jwtSecret)

	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

func ParseToken(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}
