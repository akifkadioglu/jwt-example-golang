package main

import (
	"fmt"
	"jwtExample/controller"
	"jwtExample/helpers"

	uuid "github.com/satori/go.uuid"
)

func main() {
	u := uuid.NewV4()
	fmt.Println(controller.GenerateKey(u))
	fmt.Println(helpers.DecodeKey(controller.GenerateKey(u)))
}
