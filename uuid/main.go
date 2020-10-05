package main

import (
	"fmt"
	"github.com/google/uuid"
)

func main() {
	m := make(map[string]int)
	for i:=0; i<1000000; i++ {
		uuid := uuid.New()
		m[uuid.String()] += 1
	}

	fmt.Println(len(m))
}