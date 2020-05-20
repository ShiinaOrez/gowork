package main

import (
    "fmt"
    "testing"
)

func Test_Main(m *testing.M) {
    fmt.Println("==> Test Main Start")
    init_()
    fmt.Println("==> Test Main Over")
}

func Test_Utils(t *testing.T) {
    fmt.Println("==> Test Utils Start")
    log("test u, bi*ch")
    fmt.Println("==> Test Utils Over")
}
