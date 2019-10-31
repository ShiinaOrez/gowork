package main

import (
	"fmt"
)

func main() {
	mySlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	newSlice := mySlice[1:9:3]
	fmt.Println(newSlice)
}
