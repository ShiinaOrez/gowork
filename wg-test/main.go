package main

import (
    "fmt"
    "sync"
)

func main() {
    wg := sync.WaitGroup{}
    wg.Add(1)
    select {
    case wg.Wait(): {}
    default: {
        fmt.Println("Default.")
    }
    }
    fmt.Println("Over.")
    return
}
