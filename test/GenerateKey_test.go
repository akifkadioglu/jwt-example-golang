package test

import (
	"jwtExample/controller"
	"reflect"
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestGenerateKey(t *testing.T) {
	u := uuid.NewV4()
	assert.Equal(t, reflect.TypeOf(controller.GenerateKey(u)), reflect.TypeOf(""))
}
