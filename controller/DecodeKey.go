package controller

import (
	//"jwtExample/helpers"

	"github.com/golang-jwt/jwt"
)

func DecodeKey(key string) interface{} {
	token, err := jwt.Parse(key, func(token *jwt.Token) (interface{}, error) {
		return []byte("asd"), nil
	})
	if err != nil {
		panic("Failed to create token")
	}
	claims := token.Claims.(jwt.MapClaims)

	return claims
}
