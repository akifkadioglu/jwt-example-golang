package test

import (
	"jwtExample/controller"
	"testing"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
)

func TestDecodeKey(t *testing.T) {
	exampleKey := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjhlYmJmYWU4LWNlNzEtNDkzMC05Yjk5LWYxODFhNjQ5YWYwNSIsInR5cGUiOjAsImV4cCI6MTY3MzAxNDA4OX0.2BAc526iFmdHUznzEpX9nnUVZLkvaEBVwr9MaknfQx8"
	assert.Equal(t, controller.DecodeKey(exampleKey),
		jwt.MapClaims(jwt.MapClaims{"exp": 1.673014089e+09, "id": "8ebbfae8-ce71-4930-9b99-f181a649af05", "type": 0.0}),
	)
}
