package main

import "fmt"

func main() {
	var stream chan string
	for {
		select {
		case <-stream:
		default:
		}
		fmt.Println()
	}
}
