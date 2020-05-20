package main

import "fmt"


func init_() {
    fmt.Println("==> Init... Over")
}

func log(s string) {
    fmt.Printf("[INFO]: %s\n", s)
}
