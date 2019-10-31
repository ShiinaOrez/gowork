package main

import "fmt"

func main() {
	strs := []string{"CCC", "BCD", "ACE"}
	for _, str := range strs {
		fmt.Println(HashFunc(str))
	}
}

func HashFunc(s string) (value int) {
	for _, r := range s {
		value += int(r)
	}
	return
}
