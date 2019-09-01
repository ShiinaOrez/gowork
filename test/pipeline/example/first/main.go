package main

import (
    "fmt"
    "time"
)

func gen() <-chan int {
    out := make(chan int)
    go func() {
        for n:=0; n<=1000000; n++ {
            out <- n
        }
        close(out)
    }()
    return out
}

func sq(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        for n:= range in {
            out <- n*n
        }
        close(out)
    }()
    return out
}

func main() {
    t1 := time.Now()
    for n := range sq(gen()) {
        fmt.Println(n)
    }
    t2 := time.Now()
    fmt.Println(t2.Sub(t1))
}