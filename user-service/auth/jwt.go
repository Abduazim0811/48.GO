package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtkey = []byte("my_secret_key")

type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateJWT(username,role string) (string,error){
	expirationTime:=time.Now().Add(5*time.Minute)
	claims:=&Claims{
		Username: username,
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}


	token:=jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	return token.SignedString(jwtkey)
}