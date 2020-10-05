package main

import (
    "fmt"
    "strings"
)

func main() {
    s := ":"
    l := strings.Split(s, ":")
    fmt.Println(l[1:])
}
