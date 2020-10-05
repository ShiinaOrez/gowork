package main

import (
    "fmt"
)

func main() {
    defer func() {
        if e := recover(); e != nil {
            fmt.Println("recover!")
        }
    }()
    var a = 1
    a += 1
    if a < 3 {
        panic("wow")
    }
    fmt.Println(a)
    return
}
