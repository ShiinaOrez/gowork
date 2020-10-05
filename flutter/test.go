package main

import (
    "fmt"
    "os/exec"
)


func main() {
    out, _ := exec.Command("source", "~/.bashrc").Output()
    fmt.Println(string(out))
    err := exec.Command("flutter", "create", "simple").Run()
    fmt.Println(err)
}
