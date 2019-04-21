package main

import "fmt"

func main() {
    str := "I am a string."
    for index, ru := range str {
        bytes := []byte(str)
        ints := []int(bytes)
        fmt.Println("Index:", index, "Rune:", ru)
    }
}
