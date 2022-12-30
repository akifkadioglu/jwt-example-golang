package controller

import (
	//"jwtExample/helpers"
	"jwtExample/models"
	"time"

	"github.com/golang-jwt/jwt"
	uuid "github.com/satori/go.uuid"
)

func GenerateKey(id uuid.UUID) string {
	day := time.Hour * 24
	var claims = &models.JWTModel{
		UUID: id,
		Type: 0,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(day * 7).Unix(), // For 7 day permission
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("asd"))
	if err != nil {
		panic("Failed to create token")
	}

	return t
}
