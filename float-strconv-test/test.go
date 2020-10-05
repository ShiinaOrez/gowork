package main

import (
    "fmt"
)

func main() {
    defer fmt.Println("defer~")
    panic(nil)
}

