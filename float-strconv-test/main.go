package main

import (
    "fmt"
    "strconv"
)

func main() {
    _, err := strconv.ParseFloat("", 64)
    if err != nil {
        fmt.Println("err")
    }
}
