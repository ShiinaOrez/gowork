package main

import (
	"fmt"
	"gowork/muxibook/models"
)

func main() {
	var a models.User
	a.Username = "Shiina"
	fmt.Println(a.Username)
}
