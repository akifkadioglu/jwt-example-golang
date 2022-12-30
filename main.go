package main

import (
	"fmt"
	"jwtExample/controller"

	uuid "github.com/satori/go.uuid"
)

func main() {
	u := uuid.NewV4()
	fmt.Println(controller.GenerateKey(u))
	fmt.Println(controller.DecodeKey(controller.GenerateKey(u)))
}
