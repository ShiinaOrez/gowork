package main

import (
    "fmt"
    "time"
)

func main() {
    var a [10]int
    for i:=0; i<10; i++ {
        go func(i int) {
            for {
                a[i]++
//                fmt.Println(i)
            }
        }(i)
    }
    fmt.Println("Start to sleep.")
    time.Sleep(1000)
    fmt.Println("Sleep Over.")
    fmt.Println(a)
}
