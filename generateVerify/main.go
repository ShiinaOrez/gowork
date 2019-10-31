package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func main() {
	baseNum := 10000
	num := baseNum + rand.Intn(90000)
	str, err := strconv.Itoa(num)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(str)
}
