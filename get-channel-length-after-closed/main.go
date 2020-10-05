package main

import "fmt"

func main() {
    ch := make(chan int, 5)
    for i:=0; i<3; i++ { ch<- i }
    close(ch)
    fmt.Println(len(ch))
}
