package main

import "fmt"

const mod = 1000000007

func main() {
    var a int
    fmt.Scanf("%d", &a)
    for i:=0; i<a; i++ {
        var aa int
        fmt.Scanf("%d", &aa)
    }
    
    cnt := 1
    if a == 0 {
        fmt.Printf("0")
    }
    if a == 1 {
        fmt.Printf("0")
    }
    for i:=1; i<a; i++ {
        cnt = (cnt * i) % mod
    }
    fmt.Printf("%d", cnt)
}
