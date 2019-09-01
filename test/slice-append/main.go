package main

import (
    "fmt"
    "time"
)

func main() {
    origin := []int{}
    t1 := time.Now()
    for i:=0; i<1000000; i++ {
        origin = append(origin, i)
    }
    fmt.Println(time.Now().Sub(t1))
    test := origin[:0]
    t2 := time.Now()
    for i:=0; i<1000000; i++ {
        test = append(test, i)
    }
    fmt.Println(time.Now().Sub(t2))
}
