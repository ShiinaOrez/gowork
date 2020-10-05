package main

import "fmt"

func main() {
    ch := make(chan int, 3)
    ch<- 1
    ch<- 2
    ch<- 3
    fmt.Println(len(ch))
    fmt.Println(<-ch)
    close(ch)
}
