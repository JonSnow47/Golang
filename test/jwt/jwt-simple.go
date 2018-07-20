package jwt

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const tokenKey = "SECRET"

func NewJWT(id string) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["admin"] = id
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	claims["iat"] = time.Now().Unix()
	token.Claims = claims
	token.SignedString(tokenKey)
	fmt.Println(token.Raw, token.Claims, token.Header, token.Signature, token.Valid)
}
