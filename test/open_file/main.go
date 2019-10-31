package main

import (
	"fmt"
	"log"
	"os"
)

func openFile(name string) (*os.File, error) {
	return os.Open(name)
}

func main() {
	for i := 0; i < 20000; i++ {
		_, err := openFile("a.txt")
		if err != nil {
			log.Println(err)
			fmt.Println("Open:", i-1)
			break
		}
	}
}
