// +build !test

package main

import (
	"fmt"
	"github.com/ShiinaOrez/gowork/coverage/utils"
)

func main() {
	log(utils.GetTimeJSON())
}

func log(s string) {
	fmt.Println("[ LOG ]", s)
}