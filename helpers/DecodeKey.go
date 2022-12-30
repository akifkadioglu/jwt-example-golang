package helpers

import (
	"github.com/golang-jwt/jwt"
)

func DecodeKey(key string) interface{} {
	token, _ := jwt.Parse(key, func(token *jwt.Token) (interface{}, error) {
		return []byte(GetEnvVariable("KEY")), nil
	})

	claims := token.Claims.(jwt.MapClaims)

	return claims
}
