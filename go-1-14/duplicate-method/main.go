package main

import (
    "fmt"
)

type IA interface {
    Print()
}

type IB interface {
    Print()
}

type IC interface {
    IA
    IB
}

type A struct {}
type B struct {}

func (a A) Print() { fmt.Println("A") }
func (b B) Print() { fmt.Println("B") }

func main() {
    var a IA = A{}
    var b IB = B{}
    var c IC = A{}
    var d IC = B{}

    c.Print()
    d.Print()
}
