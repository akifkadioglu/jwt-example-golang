package models

import (
	"github.com/golang-jwt/jwt"
	"github.com/satori/go.uuid"
)

type JWTModel struct {
	UUID uuid.UUID `json:"id"`
	Type int       `json:"type"`
	jwt.StandardClaims
}