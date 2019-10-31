package main

import "fmt"

func main() {
	stream := make(chan int, 10)
	for i := 0; i < 5; i++ {
		stream <- i
	}
	close(stream)
	for i := range stream {
		fmt.Println(i)
	}
}
