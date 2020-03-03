package main

import "fmt"

func genChannel() <-chan int {
    ch := make(chan int, 5)
    fmt.Println(&ch)
    defer close(ch)
    return ch
}

func main() {
    ch := genChannel()
    fmt.Println(&ch)
    for i:=0; i<5; i++ {
        fmt.Println(<-ch)
    }
}
